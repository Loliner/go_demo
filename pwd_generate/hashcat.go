package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const hashMode = "22000"

type outcome int

const (
	outcomeCracked outcome = iota
	outcomeExhausted
	outcomeAborted // 用户中断 / 运行时中断：进度已存进 restore，下次续跑
	outcomeError
)

type hashcatConfig struct {
	bin      string
	workload int
	optimize bool
	extra    []string
}

func newHashcatConfig(workload int, optimize bool, extra []string) (*hashcatConfig, error) {
	bin, err := exec.LookPath("hashcat")
	if err != nil {
		return nil, fmt.Errorf("找不到 hashcat，请确认已安装并在 PATH 中：%w", err)
	}
	return &hashcatConfig{bin: bin, workload: workload, optimize: optimize, extra: extra}, nil
}

var sessionUnsafe = regexp.MustCompile(`[^A-Za-z0-9_.-]+`)

func sessionName(pkg, wordlist string) string {
	return sessionUnsafe.ReplaceAllString(pkg+"__"+wordlist, "_")
}

// run 对一个 (原始包, 词典) 组合跑一次 hashcat。有未完成的 restore 文件就续跑。
func (c *hashcatConfig) run(pkgPath, wordlistPath, workDir string) (outcome, string, error) {
	session := sessionName(filepath.Base(pkgPath), filepath.Base(wordlistPath))
	restorePath := filepath.Join(workDir, "sessions", session+".restore")
	outfilePath := filepath.Join(workDir, "cracked", filepath.Base(pkgPath)+".out")
	if err := os.MkdirAll(filepath.Dir(restorePath), 0o755); err != nil {
		return outcomeError, "", err
	}
	if err := os.MkdirAll(filepath.Dir(outfilePath), 0o755); err != nil {
		return outcomeError, "", err
	}

	var args []string
	if _, err := os.Stat(restorePath); err == nil {
		// 断点续跑：hashcat 从 restore 文件恢复全部原始参数
		args = []string{"--session=" + session, "--restore-file-path=" + restorePath, "--restore"}
	} else {
		args = []string{
			"-m", hashMode,
			"-w", fmt.Sprintf("%d", c.workload),
			"--session=" + session,
			"--restore-file-path=" + restorePath,
			"--outfile=" + outfilePath,
			"--outfile-format=2", // 只写明文密码
			"--status",
		}
		if c.optimize {
			args = append(args, "-O")
		}
		args = append(args, c.extra...)
		args = append(args, pkgPath, wordlistPath)
	}

	cmd := exec.Command(c.bin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	code := 0
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			code = ee.ExitCode()
		} else {
			return outcomeError, "", err
		}
	}

	switch code {
	case 0:
		pw := readPassword(outfilePath)
		if pw == "" {
			pw = c.showFromPotfile(pkgPath)
		}
		return outcomeCracked, pw, nil
	case 1:
		return outcomeExhausted, "", nil
	default:
		// 2=checkpoint 3=runtime 4/255=用户退出 等：视为中断，保留 restore 下次续跑
		return outcomeAborted, "", nil
	}
}

func readPassword(outfilePath string) string {
	data, err := os.ReadFile(outfilePath)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		if line = strings.TrimRight(line, "\r"); line != "" {
			return line
		}
	}
	return ""
}

// showFromPotfile 处理 hashcat 在 potfile 里已有结果、本次没写 outfile 的情况。
func (c *hashcatConfig) showFromPotfile(pkgPath string) string {
	cmd := exec.Command(c.bin, "-m", hashMode, "--show", "--outfile-format=2", pkgPath)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return ""
	}
	for _, line := range strings.Split(buf.String(), "\n") {
		if line = strings.TrimRight(line, "\r"); line != "" {
			return line
		}
	}
	return ""
}
