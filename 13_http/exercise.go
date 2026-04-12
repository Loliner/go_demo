package main

import "fmt"

// ============================================================
// 练习：实现一个简单的 TODO REST API
// ============================================================
//
// 1. 定义 Todo struct：
//    { ID int, Title string, Done bool }
//    加上合适的 json tag
//
// 2. 用一个全局 map[int]Todo 模拟数据库
//
// 3. 实现以下路由（注册到 http.HandleFunc）：
//
//    GET  /todos        → 返回所有 todo 的 JSON 数组
//    POST /todos        → 从 body 读取 { "title": "..." }，创建新 todo，返回创建的 todo
//
// 4. 启动服务器监听 :8080
//
// 测试：
//    curl http://localhost:8080/todos
//    curl -X POST http://localhost:8080/todos -d '{"title":"Buy milk"}' -H 'Content-Type: application/json'
//    curl http://localhost:8080/todos

// TODO: 在这里写你的代码

func exercise() {
	fmt.Println("=== Exercise: TODO REST API ===")
}
