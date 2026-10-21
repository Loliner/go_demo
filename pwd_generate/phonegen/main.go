package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func padZero(num int, length int) string {
	s := strconv.Itoa(num)

	if len(s) < length {
		return strings.Repeat("0", length-len(s)) + s
	}
	return s
}

func main() {
	dir, _ := os.Getwd()
	p := filepath.Join(dir, "pwd_generate", "phone.txt")
	file, err := os.Create(p)
	if err != nil {
		fmt.Println("file created", file)
	}

	for i := 0; i <= 99999999; i++ {
		file.WriteString("4" + padZero(i, 8) + "\n")
	}
}
