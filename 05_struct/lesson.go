package main

import "fmt"

// ============================================================
// 1. 定义 struct
// 类比 JS 的 class 或 Python 的 dataclass
// ============================================================

type User struct {
	Name  string
	Age   int
	Email string
}

// ============================================================
// 2. 方法（Method）
// Go 没有 class，方法通过"接收者"绑定到 struct
// ============================================================

// 值接收者：拿到的是 struct 的副本，不能修改原始数据
func (u User) SetGreet() string {
	u.Name = "Modified" // 修改副本，不影响原始 struct
	return "Hello, " + u.Name
}

func (u *User) Greet() string {
	return "Hello, " + u.Name
}

// 指针接收者：拿到的是指针，可以修改原始数据
func (u *User) Birthday() {
	u.Age++
}

// ============================================================
// 3. 嵌套 struct（Embedding）
// ============================================================

type Address struct {
	Name    string
	City    string
	Country string
}

type Employee struct {
	User                  // 匿名嵌入，字段和方法自动提升
	Department string
	Address    Address    // 具名嵌入，需要用 .Address.City 访问
}

// ============================================================
// 4. 构造函数约定
// ============================================================

func NewUser(name string, age int, email string) *User {
	return &User{Name: name, Age: age, Email: email}
}

func lesson() {
	fmt.Println("=== 基本使用 ===")
	u1 := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	fmt.Println(u1.Name, u1.Age)

	var u2 User
	fmt.Println(u2.Name, u2.Age) // "" 0

	fmt.Println("\n=== 值类型陷阱 ===")
	a := User{Name: "Alice", Age: 30}
	b := a
	b.Name = "Bob"
	fmt.Println("a.Name:", a.Name) // Alice
	fmt.Println("b.Name:", b.Name) // Bob

	fmt.Println("\n=== 方法 ===")
	u3 := User{Name: "Charlie", Age: 25}
	fmt.Println(u3.SetGreet()) // Hello, Modified
	fmt.Println(u3.Greet())    // Hello, Charlie
	u3.Birthday()
	fmt.Println("Age after birthday:", u3.Age) // 26

	fmt.Println("\n=== 嵌套 Struct ===")
	emp := Employee{
		User:       User{Name: "Dave", Age: 28, Email: "dave@example.com"},
		Department: "Engineering",
		Address:    Address{Name: "Hi", City: "Beijing", Country: "China"},
	}
	fmt.Println(emp.Name)         // Dave（匿名嵌入，字段提升）
	fmt.Println(emp.Address.Name) // Hi（具名嵌入，需要完整路径）
	fmt.Println(emp.Greet())      // Hello, Dave
	fmt.Println(emp.Address.City) // Beijing

	fmt.Println("\n=== 构造函数 ===")
	u4 := NewUser("Eve", 22, "eve@example.com")
	fmt.Println(u4.Name, u4.Age)
	u4.Birthday()
	fmt.Println("Age after birthday:", u4.Age) // 23
}
