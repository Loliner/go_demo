package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var database = map[int]Todo{}

type Request struct {
	Title string `json:"title"`
}

var nextID int

func newID() int {
	nextID++
	return nextID
}

func getTodos() (todos []Todo) {
	for _, value := range database {
		todos = append(todos, value)
	}
	if todos == nil {
		todos = []Todo{}
	}
	return todos
}

func setTodo(title string) (todos []Todo) {
	todo := Todo{
		ID:    newID(),
		Title: title,
		Done:  false,
	}

	database[todo.ID] = todo
	fmt.Println(database)

	todos = append(todos, todo)
	return todos
}

func todoHandler(response http.ResponseWriter, request *http.Request) {
	todos := []Todo{}
	switch request.Method {
	case http.MethodGet:
		todos = getTodos()
	case http.MethodPost:
		var req Request
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			http.Error(response, "invalid json", http.StatusBadRequest)
			return
		}
		todos = setTodo(req.Title)
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(todos)
}

func exercise() {
	fmt.Println("=== Exercise: TODO REST API ===")

	http.HandleFunc("/todos", todoHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
