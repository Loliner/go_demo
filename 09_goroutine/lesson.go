package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// Go Goroutine + Channel 核心概念
// ============================================================
//
// JS 类比：
//   - goroutine ≈ async function，但更轻量（几 KB vs MB 级线程）
//   - channel ≈ 带阻塞能力的消息队列
//   - go 关键字 ≈ 你不需要写 async/await，直接 go 就跑了
//
// 核心思想（Go 官方口号）：
//   "Don't communicate by sharing memory; share memory by communicating."
//   → 不要用共享变量传数据，用 channel 传数据

// ============================================================
// 1. goroutine 基础
//    JS 里 async function 要配合 await 才能等待结果
//    Go 里 go 关键字直接让函数在新 goroutine 里跑，完全不阻塞
// ============================================================

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func demoGoroutine() {
	// 直接调用（同步，在当前 goroutine 里跑）
	sayHello("main")

	// go 关键字：在新 goroutine 里跑（异步）
	go sayHello("goroutine1")
	go sayHello("goroutine2")

	// 问题：main 函数结束后，程序退出，goroutine 可能还没跑完
	// 临时方案：Sleep 等一会儿（实际项目不这么做，用 WaitGroup 或 channel）
	time.Sleep(10 * time.Millisecond)
}

// ============================================================
// 2. WaitGroup —— 等待一组 goroutine 完成
//    类比 JS 的 Promise.all
// ============================================================

func demoWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 告诉 wg：又多了一个任务
		go func(id int) {
			defer wg.Done() // 任务完成，计数 -1
			fmt.Printf("worker %d done\n", id)
		}(i) // 注意：i 要作为参数传进去，避免闭包捕获问题（和 JS 一样的坑）
	}

	wg.Wait() // 阻塞直到所有任务完成，类比 await Promise.all(...)
	fmt.Println("all workers done")
}

// ============================================================
// 3. Channel —— goroutine 之间的通信管道
//    类比 JS 的 MessageChannel 或 Stream，但自带阻塞/同步语义
//
//    make(chan T)       → 无缓冲 channel（发送方会阻塞直到有人接收）
//    make(chan T, n)    → 有缓冲 channel（最多缓冲 n 个，满了才阻塞）
// ============================================================

func demoChannel() {
	// 无缓冲 channel：发和收必须同时准备好
	ch := make(chan int)

	go func() {
		ch <- 42 // 发送（会阻塞直到有人接收）
	}()

	val := <-ch // 接收（会阻塞直到有数据）
	fmt.Printf("received: %d\n", val)
}

// ============================================================
// 4. 用 channel 替代 WaitGroup 收集结果
//    类比 JS 的 Promise.all 但可以拿到每个结果
// ============================================================

func worker(id int, results chan<- string) {
	// chan<- 表示只能发送（send-only channel），是类型约束
	result := fmt.Sprintf("worker %d result", id)
	results <- result
}

func demoChannelResults() {
	results := make(chan string, 3) // 有缓冲，不会阻塞

	for i := 1; i <= 3; i++ {
		go worker(i, results)
	}

	// 收集 3 个结果
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
}

// ============================================================
// 5. range over channel —— 持续读取直到 channel 关闭
//    类比 JS 的 async iterator / for await...of
// ============================================================

func generate(nums []int) <-chan int {
	// <-chan 表示只能接收（receive-only channel）
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // 发完要 close，否则 range 会永久阻塞
	}()
	return out
}

func demoRangeChannel() {
	ch := generate([]int{1, 2, 3, 4, 5})

	for v := range ch { // 自动在 channel 关闭时退出
		fmt.Printf("got: %d\n", v)
	}
}

// ============================================================
// 6. select —— 同时监听多个 channel
//    类比 JS 的 Promise.race，但更灵活
// ============================================================

func demoSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		ch2 <- "two"
	}()

	// select 会选择第一个准备好的 case
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("received from ch2:", msg)
		}
	}
}

// ============================================================
// 7. done channel —— 取消/超时模式
//    类比 JS 的 AbortController
// ============================================================

func demoTimeout() {
	result := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond) // 模拟耗时操作
		result <- 42
	}()

	select {
	case v := <-result:
		fmt.Printf("got result: %d\n", v)
	case <-time.After(100 * time.Millisecond): // 超时
		fmt.Println("timeout!")
	}
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. goroutine 基础 ===")
	demoGoroutine()

	fmt.Println("\n=== 2. WaitGroup ===")
	demoWaitGroup()

	fmt.Println("\n=== 3. Channel 基础 ===")
	demoChannel()

	fmt.Println("\n=== 4. Channel 收集结果 ===")
	demoChannelResults()

	fmt.Println("\n=== 5. range over channel ===")
	demoRangeChannel()

	fmt.Println("\n=== 6. select ===")
	demoSelect()

	fmt.Println("\n=== 7. 超时模式 ===")
	demoTimeout()
}
