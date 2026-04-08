package main

import "fmt"

// ============================================================
// 练习：用 goroutine + channel 实现并发任务处理器
// ============================================================
//
// 目标：模拟一个"并发下载器"
//
// 1. 定义 Task struct：{ ID int, URL string }
//    定义 Result struct：{ TaskID int, Content string, Err error }
//
// 2. 实现 process(task Task) Result 函数，模拟处理任务：
//    - 打印 "processing task <ID>..."
//    - 假装处理成功，返回 Result{ TaskID: task.ID, Content: "data from " + task.URL }
//
// 3. 实现 runWorkers(tasks []Task, workerCount int) []Result：
//    - 创建一个 taskCh chan Task 和一个 resultCh chan Result
//    - 启动 workerCount 个 goroutine，每个从 taskCh 读任务，调用 process，把结果写入 resultCh
//    - 把所有 tasks 发送到 taskCh，发完 close(taskCh)
//    - 收集所有结果并返回
//    提示：用 WaitGroup 知道所有 worker 都完成了，然后 close(resultCh)
//
// 4. 在 exercise() 里调用 runWorkers，打印所有结果

// TODO: 在这里写你的代码

func exercise() {
	fmt.Println("=== Exercise: goroutine + channel ===")

	// 测试用例（完成代码后取消注释）:
	// tasks := []Task{
	// 	{ID: 1, URL: "https://a.com"},
	// 	{ID: 2, URL: "https://b.com"},
	// 	{ID: 3, URL: "https://c.com"},
	// 	{ID: 4, URL: "https://d.com"},
	// 	{ID: 5, URL: "https://e.com"},
	// }
	// results := runWorkers(tasks, 3)
	// for _, r := range results {
	// 	if r.Err != nil {
	// 		fmt.Printf("task %d failed: %v\n", r.TaskID, r.Err)
	// 	} else {
	// 		fmt.Printf("task %d: %s\n", r.TaskID, r.Content)
	// 	}
	// }
}
