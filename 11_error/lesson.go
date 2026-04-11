package main

import (
	"errors"
	"fmt"
)

// ============================================================
// Go error 处理模式
// ============================================================
//
// JS 类比：
//   - Go 没有 try/catch，错误是普通返回值
//   - 类比 Node.js 的 callback(err, result) 风格，但用多返回值实现
//   - error 是一个 interface，不是特殊类型

// ============================================================
// 1. error 是 interface
//    内置定义只有一个方法：Error() string
//    类比 JS 的 Error 对象，但任何实现了 Error() string 的类型都是 error
// ============================================================

// error interface 的定义（Go 内置，不需要你写）：
// type error interface {
//     Error() string
// }

func demoErrorInterface() {
	// errors.New 创建最简单的 error
	err := errors.New("something went wrong")
	fmt.Println(err)         // 调用 err.Error()
	fmt.Println(err.Error()) // 显式调用

	// fmt.Errorf 支持格式化，类比 JS 的 new Error(`msg: ${detail}`)
	name := "config.json"
	err2 := fmt.Errorf("file not found: %s", name)
	fmt.Println(err2)
}

// ============================================================
// 2. 函数返回 error —— Go 最常见的错误处理方式
//    惯例：error 放最后一个返回值，成功时返回 nil
// ============================================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero: %v / %v", a, b)
	}
	return a / b, nil
}

func demoReturnError() {
	// 惯用写法：if err != nil { ... }
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10/2 =", result)

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// ============================================================
// 3. 自定义 error 类型
//    类比 JS 的 class ValidationError extends Error
//    好处：调用方可以用 errors.As 判断具体类型，拿到额外字段
// ============================================================

type ValidationError struct {
	Field   string
	Message string
}

// 实现 error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %q: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

func demoCustomError() {
	err := validateAge(-1)
	if err != nil {
		fmt.Println(err) // 调用 Error() 方法

		// errors.As：类比 JS 的 err instanceof ValidationError
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("field: %s, message: %s\n", ve.Field, ve.Message)
		}
	}
}

// ============================================================
// 4. error wrapping —— 给 error 加上调用链上下文
//    类比 JS 的 new Error("context", { cause: originalError })
//    Go 1.13+ 支持 %w 动词
// ============================================================

func readConfig(path string) error {
	// 模拟底层错误
	err := fmt.Errorf("permission denied")
	// %w 包装：保留原始 error，同时加上上下文
	return fmt.Errorf("readConfig(%q): %w", path, err)
}

func loadApp() error {
	err := readConfig("/etc/app.conf")
	if err != nil {
		return fmt.Errorf("loadApp: %w", err)
	}
	return nil
}

func demoWrapping() {
	err := loadApp()
	if err != nil {
		fmt.Println(err) // 完整链：loadApp: readConfig(...): permission denied

		// 反例：用相同文字新建一个 error，errors.Is 仍然返回 false
		// errors.Is 比较的是对象引用（类比 JS 的 ===），不是字符串内容
		target := fmt.Errorf("permission denied")                    // 新实例，和链里的不是同一个对象
		fmt.Println("is permission denied?", errors.Is(err, target)) // false

		// 实际生产中通常用 sentinel error 配合 errors.Is
	}
}

// ============================================================
// 5. sentinel error —— 预定义的"哨兵"错误值
//    类比 JS 的 export const ERR_NOT_FOUND = new Error("not found")
//    用 errors.Is 比较，不要用 ==（因为 wrapping 后用 == 会失效）
// ============================================================

var ErrNotFound = errors.New("not found")
var ErrPermission = errors.New("permission denied")

func findUser(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	if id < 0 {
		return "", fmt.Errorf("findUser(%d): %w", id, ErrPermission) // 包装 sentinel
	}
	return fmt.Sprintf("user_%d", id), nil
}

func demoSentinel() {
	_, err := findUser(0)
	fmt.Println("errors.Is ErrNotFound:", errors.Is(err, ErrNotFound)) // true

	_, err = findUser(-1)
	fmt.Println("errors.Is ErrPermission:", errors.Is(err, ErrPermission)) // true，即使被 wrap 了
	fmt.Println("== ErrPermission:", err == ErrPermission)                 // false！wrap 后 == 失效
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. error interface ===")
	demoErrorInterface()

	fmt.Println("\n=== 2. 函数返回 error ===")
	demoReturnError()

	fmt.Println("\n=== 3. 自定义 error 类型 ===")
	demoCustomError()

	fmt.Println("\n=== 4. error wrapping ===")
	demoWrapping()

	fmt.Println("\n=== 5. sentinel error ===")
	demoSentinel()
}
