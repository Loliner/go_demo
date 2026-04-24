package main

import (
	"context"
	"fmt"
	"time"
)

// ============================================================
// 练习：带超时和取消的任务调度器
// ============================================================
//
// 实现函数 runWithTimeout：
//
//   func runWithTimeout(timeout time.Duration, task func(ctx context.Context) error) error
//
//   - 用 context.WithTimeout 创建一个带超时的 context
//   - 在 goroutine 中运行 task(ctx)
//   - 如果 task 在超时前完成，返回 task 的 error
//   - 如果超时，返回 context.DeadlineExceeded
//
// 提示：
//   - 用 channel 在 goroutine 和主函数之间传递 task 的结果
//   - 用 select 同时等待 task 完成 和 ctx.Done()

// TODO: 在这里写你的代码

func runWithTimeout(timeout time.Duration, task func(ctx context.Context) error) error {
	// panic("not implemented")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultCh := make(chan error, 1)

	go func() {
		resultCh <- task(ctx)
	}()

	select {
	case result := <-resultCh:
		return result
	case <-ctx.Done():
		return ctx.Err()
	}
}

func exercise() {
	fmt.Println("=== Exercise: runWithTimeout ===")

	// 测试 1：task 在超时前完成
	err := runWithTimeout(1*time.Second, func(ctx context.Context) error {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("task 1 done")
		return nil
	})
	fmt.Println("result 1:", err) // <nil>

	// 测试 2：task 超时
	err = runWithTimeout(200*time.Millisecond, func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		fmt.Println("task 2 done") // 不应该打印
		return nil
	})
	fmt.Println("result 2:", err) // context deadline exceeded

	// 测试 3：task 主动响应取消，提前退出
	err = runWithTimeout(200*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("task 3 done") // 不应该打印
			return nil
		case <-ctx.Done():
			fmt.Println("task 3 canceled")
			return ctx.Err()
		}
	})
	fmt.Println("result 3:", err) // context deadline exceeded
}
