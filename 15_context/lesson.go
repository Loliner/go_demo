package main

import (
	"context"
	"fmt"
	"time"
)

// ============================================================
// Go context 包
// ============================================================
//
// JS 类比：
//   context.WithCancel   ≈ AbortController
//   context.WithTimeout  ≈ AbortController + setTimeout
//   context.WithValue    ≈ AsyncLocalStorage（请求级别的隐式传参）
//
// 核心接口：
//   type Context interface {
//       Done() <-chan struct{}   // 被取消时关闭，类比 signal.aborted 变为 true
//       Err()  error            // 取消原因：Canceled 或 DeadlineExceeded
//       Deadline() (time.Time, bool)
//       Value(key any) any
//   }

// ============================================================
// 1. context.WithCancel —— 手动取消
// ============================================================

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// ctx.Done() 是一个 channel，取消时被关闭
			fmt.Printf("worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("worker %d working...\n", id)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func demoCancel() {
	// context.Background() 是根 context，类比 AbortController 的 signal 还没有 abort
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(700 * time.Millisecond)

	cancel() // 触发取消，类比 controller.abort()
	// cancel 可以安全地多次调用

	time.Sleep(500 * time.Millisecond) // 等 goroutine 打印完
}

// ============================================================
// 2. context.WithTimeout —— 超时自动取消
// ============================================================

func fetchData(ctx context.Context) (string, error) {
	// 模拟一个耗时操作
	select {
	case <-time.After(500 * time.Millisecond): // 模拟 500ms 的工作
		return "data", nil
	case <-ctx.Done(): // 超时或被取消
		return "", ctx.Err()
	}
}

func demoTimeout() {
	// 设置 200ms 超时，fetchData 需要 500ms → 必然超时
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // 即使提前完成也要 cancel，释放资源（好习惯）

	result, err := fetchData(ctx)
	if err != nil {
		fmt.Println("timeout error:", err) // context deadline exceeded
		return
	}
	fmt.Println("result:", result)
}

func demoTimeoutOK() {
	// 设置 1s 超时，fetchData 需要 500ms → 正常完成
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := fetchData(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result) // result: data
}

// ============================================================
// 3. context.WithValue —— 传递请求级别的数据
// ============================================================
//
// 类比 Node.js 的 AsyncLocalStorage，用于传递 request-scoped 数据
// （如 request ID、用户信息），不需要层层传参
//
// 注意：key 要用自定义类型，避免不同包的 key 冲突

type contextKey string

const requestIDKey contextKey = "requestID"

func handleRequest(ctx context.Context) {
	// 从 context 取值，类比 asyncLocalStorage.getStore()
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("handling request:", id)
	}
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("processing request:", id)
	}
}

func demoValue() {
	ctx := context.WithValue(context.Background(), requestIDKey, "req-abc-123")
	handleRequest(ctx)
}

// ============================================================
// 4. context 在 HTTP handler 中的使用（最常见场景）
// ============================================================
//
// net/http 的每个请求自带 context，通过 r.Context() 获取。
// 用户断开连接时，这个 context 会自动取消。
//
// func myHandler(w http.ResponseWriter, r *http.Request) {
//     ctx := r.Context()
//     result, err := db.QueryContext(ctx, "SELECT ...")  // 传给数据库查询
//     if err != nil {
//         // 用户断开 → context canceled → 查询自动取消
//     }
// }

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. WithCancel ===")
	demoCancel()

	fmt.Println("\n=== 2. WithTimeout (超时) ===")
	demoTimeout()

	fmt.Println("\n=== 2. WithTimeout (正常) ===")
	demoTimeoutOK()

	fmt.Println("\n=== 3. WithValue ===")
	demoValue()
}
