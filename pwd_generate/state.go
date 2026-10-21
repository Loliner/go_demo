package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	statusPending   = "pending"
	statusExhausted = "exhausted"
	statusCracked   = "cracked"
	statusMissing   = "missing"
)

type packageState struct {
	Cracked     bool              `json:"cracked"`
	Password    string            `json:"password,omitempty"`
	CrackedWith string            `json:"crackedWith,omitempty"`
	Wordlists   map[string]string `json:"wordlists"`
}

func (p *packageState) status(wl string) string {
	if s, ok := p.Wordlists[wl]; ok {
		return s
	}
	return statusPending
}

type state struct {
	Packages       map[string]*packageState `json:"packages"`
	KnownWordlists []string                 `json:"knownWordlists"`
	Review         []string                 `json:"review,omitempty"`
}

func loadState(path string) (*state, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &state{Packages: map[string]*packageState{}}, nil
	}
	if err != nil {
		return nil, err
	}
	var s state
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	if s.Packages == nil {
		s.Packages = map[string]*packageState{}
	}
	return &s, nil
}

func (s *state) save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func (s *state) knownWordlist(wl string) bool {
	for _, k := range s.KnownWordlists {
		if k == wl {
			return true
		}
	}
	return false
}

// renderMarkdown 把 state 渲染成给人看的 record.md。wordlists 是当前磁盘上实际存在的
// 词典(有序)，用来标出每个包还有哪些没跑；state 里记过但现在不在磁盘上的词典单独标注。
func (s *state) renderMarkdown(path string, pkgs, wordlists []string) error {
	var b strings.Builder
	b.WriteString("# WiFi 破解进度\n\n")
	b.WriteString("> 由 hashcrack 自动生成，请勿手改（下次运行会被覆盖）。\n")
	fmt.Fprintf(&b, "> 最后更新：%s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	cracked, running := 0, 0
	for _, pkg := range pkgs {
		if ps := s.Packages[pkg]; ps != nil && ps.Cracked {
			cracked++
		} else {
			running++
		}
	}
	fmt.Fprintf(&b, "共 %d 个原始包：已破解 %d，进行中 %d。\n\n", len(pkgs), cracked, running)

	for _, pkg := range pkgs {
		ps := s.Packages[pkg]
		if ps == nil {
			ps = &packageState{}
		}
		if ps.Cracked {
			fmt.Fprintf(&b, "## ✅ %s\n\n", pkg)
			fmt.Fprintf(&b, "- 密码：`%s`（命中词典：%s）\n\n", ps.Password, ps.CrackedWith)
			continue
		}
		fmt.Fprintf(&b, "## ⏳ %s\n\n", pkg)
		seen := map[string]bool{}
		for _, wl := range wordlists {
			seen[wl] = true
			fmt.Fprintf(&b, "- %s — %s\n", wl, statusLabel(ps.status(wl)))
		}
		// state 里记过但磁盘上已经没有的词典
		extra := make([]string, 0)
		for wl := range ps.Wordlists {
			if !seen[wl] {
				extra = append(extra, wl)
			}
		}
		sort.Strings(extra)
		for _, wl := range extra {
			fmt.Fprintf(&b, "- %s — %s（文件已不在磁盘）\n", wl, statusLabel(ps.status(wl)))
		}
		b.WriteString("\n")
	}

	if len(s.Review) > 0 {
		b.WriteString("## 待人工核对\n\n")
		b.WriteString("> 从旧 record.txt 迁移时无法确定的记录，需人工确认。\n\n")
		for _, r := range s.Review {
			fmt.Fprintf(&b, "- %s\n", r)
		}
		b.WriteString("\n")
	}

	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func statusLabel(s string) string {
	switch s {
	case statusCracked:
		return "✅ 已命中"
	case statusExhausted:
		return "跑完未命中"
	case statusMissing:
		return "⚠️ 词典缺失，已跳过"
	default:
		return "待跑"
	}
}
