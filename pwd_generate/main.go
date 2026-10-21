package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	root        string
	pkgDir      string
	passwordDir string
	workDir     string
	stateFile   string
	recordFile  string
	syncOnly    bool
}

func main() {
	root := flag.String("root", "", "包含 pkg/ 和 password/ 的目录，留空则自动探测")
	syncOnly := flag.Bool("sync-only", false, "只扫描登记新增，不执行破解")
	workload := flag.Int("w", 3, "hashcat 负载档位 1-4")
	optimize := flag.Bool("O", true, "启用 hashcat 优化内核 -O（限制密码最大长度）")
	extra := flag.String("extra", "", "追加给 hashcat 的额外参数，空格分隔，如 \"-d 1\"")
	flag.Parse()

	cfg := resolveConfig(*root, *syncOnly)

	if err := run(cfg, *workload, *optimize, strings.Fields(*extra)); err != nil {
		fmt.Fprintln(os.Stderr, "错误:", err)
		os.Exit(1)
	}
}

func resolveConfig(root string, syncOnly bool) *config {
	if root == "" {
		root = "."
		if _, err := os.Stat(filepath.Join(root, "pkg")); err != nil {
			if _, err := os.Stat(filepath.Join("pwd_generate", "pkg")); err == nil {
				root = "pwd_generate"
			}
		}
	}
	work := filepath.Join(root, ".hashcrack")
	return &config{
		root:        root,
		pkgDir:      filepath.Join(root, "pkg"),
		passwordDir: filepath.Join(root, "password"),
		workDir:     work,
		stateFile:   filepath.Join(work, "state.json"),
		recordFile:  filepath.Join(root, "record.md"),
		syncOnly:    syncOnly,
	}
}

func run(cfg *config, workload int, optimize bool, extra []string) error {
	st, err := loadState(cfg.stateFile)
	if err != nil {
		return fmt.Errorf("读取 state.json: %w", err)
	}

	pkgs, err := listPackages(cfg.pkgDir)
	if err != nil {
		return fmt.Errorf("扫描 pkg/: %w", err)
	}
	wordlists, err := listWordlists(cfg.passwordDir)
	if err != nil {
		return fmt.Errorf("扫描 password/: %w", err)
	}

	sync(st, pkgs, wordlists)

	if err := st.save(cfg.stateFile); err != nil {
		return err
	}
	if err := st.renderMarkdown(cfg.recordFile, pkgs, wordlists); err != nil {
		return err
	}

	if cfg.syncOnly {
		printTodo(st, pkgs, wordlists)
		return nil
	}

	hc, err := newHashcatConfig(workload, optimize, extra)
	if err != nil {
		return err
	}
	return crackAll(cfg, st, hc, pkgs, wordlists)
}

// sync 把新原始包、新密码本登记进 state，并提示 state 里记过但磁盘上已消失的词典。
func sync(st *state, pkgs, wordlists []string) {
	for _, pkg := range pkgs {
		if _, ok := st.Packages[pkg]; !ok {
			st.Packages[pkg] = &packageState{Wordlists: map[string]string{}}
			fmt.Printf("新增原始包: %s\n", pkg)
		}
	}
	present := map[string]bool{}
	for _, wl := range wordlists {
		present[wl] = true
		if !st.knownWordlist(wl) {
			st.KnownWordlists = append(st.KnownWordlists, wl)
			fmt.Printf("新增密码本: %s\n", wl)
		}
	}
	for _, wl := range st.KnownWordlists {
		if !present[wl] {
			fmt.Printf("提示: 密码本 %s 已登记但磁盘上不存在，破解时将跳过\n", wl)
		}
	}
}

func crackAll(cfg *config, st *state, hc *hashcatConfig, pkgs, wordlists []string) error {
	for _, pkg := range pkgs {
		ps := st.Packages[pkg]
		if ps.Wordlists == nil {
			ps.Wordlists = map[string]string{}
		}
		if ps.Cracked {
			fmt.Printf("[跳过] %s 已破解: %s\n", pkg, ps.Password)
			continue
		}

		pkgPath := filepath.Join(cfg.pkgDir, pkg)
		for _, wl := range wordlists {
			switch ps.status(wl) {
			case statusCracked, statusExhausted:
				continue
			}

			wlPath := filepath.Join(cfg.passwordDir, filepath.FromSlash(wl))
			if _, err := os.Stat(wlPath); err != nil {
				ps.Wordlists[wl] = statusMissing
				fmt.Printf("提示: 密码本 %s 不存在，跳过\n", wl)
				if err := persist(cfg, st, pkgs, wordlists); err != nil {
					return err
				}
				continue
			}

			fmt.Printf("\n==> 破解 %s × %s\n", pkg, wl)
			res, pw, err := hc.run(pkgPath, wlPath, cfg.workDir)
			if err != nil {
				return err
			}
			switch res {
			case outcomeCracked:
				ps.Cracked = true
				ps.Password = pw
				ps.CrackedWith = wl
				ps.Wordlists[wl] = statusCracked
				fmt.Printf("[命中] %s 密码: %s\n", pkg, pw)
				if err := persist(cfg, st, pkgs, wordlists); err != nil {
					return err
				}
			case outcomeExhausted:
				ps.Wordlists[wl] = statusExhausted
				if err := persist(cfg, st, pkgs, wordlists); err != nil {
					return err
				}
			case outcomeAborted:
				fmt.Println("hashcat 被中断，进度已保存，下次运行将从断点续跑。")
				return persist(cfg, st, pkgs, wordlists)
			}
			if ps.Cracked {
				break
			}
		}
	}
	fmt.Println("\n全部原始包处理完毕。")
	return nil
}

func persist(cfg *config, st *state, pkgs, wordlists []string) error {
	if err := st.save(cfg.stateFile); err != nil {
		return err
	}
	return st.renderMarkdown(cfg.recordFile, pkgs, wordlists)
}

func printTodo(st *state, pkgs, wordlists []string) {
	fmt.Println("\n待办（--sync-only）:")
	for _, pkg := range pkgs {
		ps := st.Packages[pkg]
		if ps.Cracked {
			continue
		}
		pending := 0
		for _, wl := range wordlists {
			if s := ps.status(wl); s != statusCracked && s != statusExhausted {
				pending++
			}
		}
		fmt.Printf("  %s: 待跑词典 %d/%d\n", pkg, pending, len(wordlists))
	}
}
