package main

import "fmt"

// ============================================================
// 练习：defer / panic / recover
// ============================================================
//
// 目标：实现一个带 recover 的安全执行器
//
// 1. 实现 safeRun(fn func()) (err error)：
//    - 用 defer + recover 捕获 fn() 内部的 panic
//    - 如果 panic 了，返回 fmt.Errorf("panic: %v", r)
//    - 正常执行则返回 nil
//
// 2. 实现 mustPositive(n int) int：
//    - 如果 n <= 0，panic("n must be positive")
//    - 否则返回 n
//
// 3. 在 exercise() 里：
//    - 用 safeRun 调用 mustPositive(5)，打印结果
//    - 用 safeRun 调用 mustPositive(-1)，打印捕获到的 error
//    - 打印 "program still running" 证明 recover 生效

// TODO: 在这里写你的代码

func exercise() {
	fmt.Println("=== Exercise: defer / panic / recover ===")
}
