package main

import "fmt"

// ============================================================
// Go defer / panic / recover 核心概念
// ============================================================
//
// JS 类比：
//   - defer    ≈ try { ... } finally { cleanup() }，但更轻量
//   - panic    ≈ throw new Error(...)，但会向上传播直到崩溃
//   - recover  ≈ catch，但只能在 defer 里用

// ============================================================
// 1. defer 基础 —— 延迟执行，函数退出前才跑
//    类比 JS 的 finally 块
// ============================================================

func demoDefer() {
	fmt.Println("start")
	defer fmt.Println("deferred 1") // 注册延迟，函数结束前执行
	defer fmt.Println("deferred 2")
	defer fmt.Println("deferred 3")
	fmt.Println("end")

	// 输出顺序：start → end → deferred 3 → deferred 2 → deferred 1
	// defer 是栈（LIFO）：后注册的先执行
}

// ============================================================
// 2. defer 的典型用途：资源清理
//    类比 JS 的 finally 或 using（TC39 proposal）
// ============================================================

func readFile(name string) {
	fmt.Printf("opening %s\n", name)
	// 实际场景：f, err := os.Open(name)
	defer fmt.Printf("closing %s\n", name) // 无论后面发生什么，都会执行

	fmt.Printf("reading %s\n", name)
	// 就算这里 panic，defer 也会执行
}

func demoCleanup() {
	readFile("data.txt")
}

// ============================================================
// 3. defer + 闭包 —— 捕获变量的时机
//    坑：defer 注册时只捕获值，不是引用（基本类型）
// ============================================================

func demoClosureCapture() {
	x := 0
	defer func() {
		fmt.Println("defer sees x =", x) // 捕获的是变量本身（引用），所以是最终值
	}()
	x = 100
	fmt.Println("x is now", x)

	// 输出：x is now 100 → defer sees x = 100
	// 和 JS 闭包一样：defer 里的匿名函数捕获的是变量引用，不是快照
}

func demoDeferValue() {
	// 对比：如果 defer 直接传参，参数在注册时求值（值拷贝）
	x := 0
	defer fmt.Println("defer param x =", x) // x 在这里就被求值为 0
	x = 100
	fmt.Println("x is now", x)

	// 输出：x is now 100 → defer param x = 0
	// 注意：直接传参 vs 闭包捕获，行为不同！
}

// ============================================================
// 4. panic —— 程序崩溃，类比 throw
//    panic 会：
//    1. 立即停止当前函数执行
//    2. 沿调用栈向上传播
//    3. 执行沿途所有 defer
//    4. 最终让程序崩溃并打印堆栈
// ============================================================

func riskyOp() {
	fmt.Println("before panic")
	panic("something went wrong!") // 类比 throw new Error("...")
	fmt.Println("after panic")     // 永远不会执行
}

func demoPanic() {
	defer fmt.Println("defer in demoPanic runs even on panic")
	fmt.Println("calling riskyOp")
	riskyOp()
	fmt.Println("this won't run")
}

// ============================================================
// 5. recover —— 捕获 panic，类比 catch
//    recover 只能在 defer 里调用，否则无效
//    recover 返回 panic 传入的值，没有 panic 则返回 nil
// ============================================================

func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero!")
	}
	return a / b, nil
}

func demoRecover() {
	result, err := safeDiv(10, 2)
	fmt.Printf("10/2 = %d, err = %v\n", result, err)

	result, err = safeDiv(10, 0)
	fmt.Printf("10/0 = %d, err = %v\n", result, err)

	fmt.Println("program continues after recover")
}

// ============================================================
// 6. 命名返回值 + defer —— 可以修改返回值
//    这个是 Go 独有的技巧
// ============================================================

func demoNamedReturn() (result string) {
	defer func() {
		result = result + " (modified by defer)" // 可以直接修改命名返回值
	}()

	result = "original"
	return // 裸 return，返回当前 result 值
	// 但 defer 会在 return 之后、真正退出之前运行，所以能修改
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. defer 基础（LIFO 顺序）===")
	demoDefer()

	fmt.Println("\n=== 2. defer 资源清理 ===")
	demoCleanup()

	fmt.Println("\n=== 3a. defer + 闭包捕获变量引用 ===")
	demoClosureCapture()

	fmt.Println("\n=== 3b. defer 直接传参（注册时求值）===")
	demoDeferValue()

	fmt.Println("\n=== 4. panic ===")
	// demoPanic() // 取消注释会让程序崩溃（没有 recover）

	fmt.Println("\n=== 5. recover ===")
	demoRecover()

	fmt.Println("\n=== 6. 命名返回值 + defer ===")
	fmt.Println(demoNamedReturn())
}
