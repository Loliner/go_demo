package basic

import (
	"errors"
	"fmt"
)

// --- 1. 多返回值 ---
// JS 里要返回多个值只能用对象或数组，Go 原生支持
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // nil 表示没有错误，类比 JS 的 null
}

// --- 2. error 处理：Go 的惯用法 ---
func main() {
	// Go 没有 try/catch，错误是普通的返回值
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("10 / 2 =", result)

	// 错误情况
	result2, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // 走这里
		return
	}
	fmt.Println("10 / 0 =", result2)

	// --- 3. defer ---
	// defer 注册一个函数，在当前函数返回前执行
	// 类比 JS 的 finally，但可以注册多次，按 LIFO 顺序执行
	fmt.Println("--- defer demo ---")
	deferDemo()
}

func deferDemo() {
	defer fmt.Println("defer 1: 最后执行")
	defer fmt.Println("defer 2: 倒数第二执行")

	fmt.Println("函数体执行中...")
	// 实际用途：确保资源释放，类比 JS 的 finally
	// 比如：defer file.Close()、defer db.Close()
}
