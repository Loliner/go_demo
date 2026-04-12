package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// ============================================================
// Go 标准库：encoding/json + net/http
// ============================================================
//
// JS 类比：
//   - encoding/json  ≈ JSON.stringify / JSON.parse
//   - net/http       ≈ Node.js 的 http 模块 / Express（但更底层）

// ============================================================
// 1. encoding/json —— 序列化与反序列化
// ============================================================

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`             // 忽略，不参与序列化
	Age      int    `json:"age,omitempty"` // 零值时省略
}

func demoJSON() {
	// Marshal：struct → JSON 字符串，类比 JSON.stringify
	u := User{ID: 1, Name: "Alice", Password: "secret", Age: 0}
	data, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // {"id":1,"name":"Alice"}  Password 忽略，Age 省略

	// Unmarshal：JSON 字符串 → struct，类比 JSON.parse
	raw := `{"id":2,"name":"Bob","age":30}`
	var u2 User
	if err := json.Unmarshal([]byte(raw), &u2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u2) // {ID:2 Name:Bob Password: Age:30}

	// map 也可以序列化/反序列化
	m := map[string]any{"key": "value", "num": 42}
	data2, _ := json.Marshal(m)
	fmt.Println(string(data2))
}

// ============================================================
// 2. HTTP 客户端
//    类比 fetch / axios
// ============================================================

func demoHTTPClient() {
	// GET 请求
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("GET error:", err)
		return
	}
	defer resp.Body.Close() // 必须关闭 body，否则连接泄漏

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("status:", resp.StatusCode)
	fmt.Println("body length:", len(body))

	// POST JSON
	payload := `{"name":"Alice"}`
	resp2, err := http.Post(
		"https://httpbin.org/post",
		"application/json",
		strings.NewReader(payload),
	)
	if err != nil {
		fmt.Println("POST error:", err)
		return
	}
	defer resp2.Body.Close()
	fmt.Println("POST status:", resp2.StatusCode)
}

// ============================================================
// 3. HTTP 服务器
//    类比 Node.js 的 http.createServer 或 Express
// ============================================================

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// 只接受 POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取 Header
	fmt.Printf("Headers: %+v\n", r.Header)
	fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))

	// 读取并解析 JSON body
	var req GreetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// 返回 JSON 响应
	resp := GreetResponse{Message: "Hello, " + req.Name + "!"}
	w.Header().Set("Content-Type", "application/json") // 设置返回头
	json.NewEncoder(w).Encode(resp)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w 是 ResponseWriter，类比 Express 的 res
	// r 是 *Request，类比 Express 的 req
	fmt.Fprintf(w, "Hello, World!")
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. encoding/json ===")
	demoJSON()

	// 2. HTTP 客户端
	fmt.Println("\n=== 2. HTTP 客户端 ===")
	demoHTTPClient()

	fmt.Println("\n=== 3. HTTP 服务器 ===")
	// 注册路由，类比 Express 的 app.get / app.post
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("try: curl http://localhost:8080/hello")
	fmt.Println(`try: curl -X POST http://localhost:8080/greet -d '{"name":"Alice"}' -H 'Content-Type: application/json'`)
	fmt.Println("starting server on :8080")
	// ListenAndServe 会阻塞，类比 app.listen(8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// 稍等服务器启动
	// time.Sleep(100 * time.Millisecond)

}
