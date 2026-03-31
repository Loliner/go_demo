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
// 类比 JS：function greet() { return "Hello, " + this.name }（但 this 是拷贝）
func (u User) SetGreet() string {
	u.Name = "Modified" // 修改副本，不影响原始 struct
	return "Hello, " + u.Name
}
func (u *User) Greet() string {
	return "Hello, " + u.Name
}

// 指针接收者：拿到的是指针，可以修改原始数据
// 类比 JS：正常的 class 方法（this 是引用）
func (u *User) Birthday() {
	u.Age++ // 修改原始 struct
}

// ============================================================
// 3. 嵌套 struct（Embedding）
// 类比 JS class 的继承，但 Go 用组合而不是继承
// ============================================================

type Address struct {
	Name    string
	City    string
	Country string
}

type Employee struct {
	User       // 嵌入 User（匿名字段），自动"继承"所有字段和方法
	Department string
	Address    Address // 具名嵌入，需要用 .Address.City 访问
}

// ============================================================
// 4. 构造函数约定
// Go 没有 constructor 关键字，约定用 NewXxx 函数
// 类比 JS 的 static create() 或 factory function
// ============================================================

func NewUser(name string, age int, email string) *User {
	return &User{ // & 取地址，返回指针（下一章讲）
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func main() {
	// --------------------------------------------------------
	// 基本使用
	// --------------------------------------------------------
	fmt.Println("=== 基本使用 ===")

	// 字面量创建（推荐用字段名，不怕字段顺序变化）
	u1 := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	fmt.Println(u1.Name, u1.Age) // Alice 30

	// 零值：未赋值的字段自动为零值，不会 undefined/nil panic
	var u2 User
	fmt.Println(u2.Name, u2.Age) // "" 0

	// --------------------------------------------------------
	// 值类型 vs 引用类型（重点！）
	// --------------------------------------------------------
	fmt.Println("\n=== 值类型陷阱 ===")

	// JS：let b = a 是引用拷贝，Go：b := a 是值拷贝
	a := User{Name: "Alice", Age: 30}
	b := a // 完整拷贝，b 是独立的
	b.Name = "Bob"
	fmt.Println("a.Name:", a.Name) // Alice（不受 b 影响）
	fmt.Println("b.Name:", b.Name) // Bob

	// --------------------------------------------------------
	// 方法调用
	// --------------------------------------------------------
	fmt.Println("\n=== 方法 ===")

	u3 := User{Name: "Charlie", Age: 25}
	fmt.Println(u3.SetGreet()) // Hello, Modified（值接收者，修改了副本）
	fmt.Println(u3.Greet())    // Hello, Charlie

	u3.Birthday()                              // 指针接收者：Age 从 25 变 26
	fmt.Println("Age after birthday:", u3.Age) // 26

	// --------------------------------------------------------
	// 嵌套 struct
	// --------------------------------------------------------
	fmt.Println("\n=== 嵌套 Struct ===")

	emp := Employee{
		User:       User{Name: "Dave", Age: 28, Email: "dave@example.com"},
		Department: "Engineering",
		Address:    Address{Name: "Hi", City: "Beijing", Country: "China"}, // Address 是具名嵌入，需要完整路径访问
	}

	// 匿名嵌入的字段和方法可以直接访问（提升）
	fmt.Println(emp.Name)         // Dave（等价于 emp.User.Name）
	fmt.Println(emp.Address.Name) // Hi
	fmt.Println(emp.Greet())      // Hello, Dave（等价于 emp.User.Greet()）
	fmt.Println(emp.Department)   // Engineering

	// 具名嵌入需要完整路径
	fmt.Println(emp.Address.City) // Beijing

	// --------------------------------------------------------
	// 构造函数
	// --------------------------------------------------------
	fmt.Println("\n=== 构造函数 ===")

	u4 := NewUser("Eve", 22, "eve@example.com")
	fmt.Println(u4.Name, u4.Age) // Eve 22
	u4.Birthday()
	fmt.Println("Age after birthday:", u4.Age) // 23
}
