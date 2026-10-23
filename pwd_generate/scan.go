package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var wordlistExts = map[string]bool{".txt": true, ".dic": true, ".found": true}

// listPackages 返回 pkgDir 下所有 .22000 原始包，按文件名排序。
func listPackages(pkgDir string) ([]string, error) {
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".22000") {
			out = append(out, e.Name())
		}
	}
	sort.Strings(out)
	return out, nil
}

// listWordlists 递归扫 passwordDir 下的明文词典，返回相对路径(统一用 / 分隔)。
// 顺序：先当前目录的文件(按大小从小到大)，再按名字进入各子目录递归。
func listWordlists(passwordDir string) ([]string, error) {
	return walkWordlists(passwordDir, passwordDir)
}

func walkWordlists(base, cur string) ([]string, error) {
	entries, err := os.ReadDir(cur)
	if err != nil {
		return nil, err
	}
	var files []string
	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
			continue
		}
		if !wordlistExts[strings.ToLower(filepath.Ext(e.Name()))] {
			continue
		}
		files = append(files, e.Name())
	}
	sort.Strings(files)
	sort.Strings(dirs)

	var out []string
	for _, f := range files {
		rel, err := filepath.Rel(base, filepath.Join(cur, f))
		if err != nil {
			continue
		}
		out = append(out, filepath.ToSlash(rel))
	}
	for _, d := range dirs {
		sub, err := walkWordlists(base, filepath.Join(cur, d))
		if err != nil {
			return nil, err
		}
		out = append(out, sub...)
	}
	return out, nil
}
