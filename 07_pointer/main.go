package main

import "fmt"

// ============================================================
// 1. 指针基础
// JS/Python 里你感知不到指针，Go 把它显式暴露出来
// 指针就是：存储另一个变量内存地址的变量
// ============================================================

func basicPointer() {
	x := 42
	p := &x // & 取地址：p 存的是 x 的内存地址

	fmt.Println("x =", x)   // 42
	fmt.Println("p =", p)   // 0xc000...（内存地址）
	fmt.Println("*p =", *p) // 42，* 解引用：拿到地址指向的值

	*p = 100             // 通过指针修改 x 的值
	fmt.Println("x =", x) // 100，x 被修改了
}

// ============================================================
// 2. 为什么需要指针：函数传参默认是值拷贝
// ============================================================

// 值传递：拿到的是副本，修改不影响原始值
// 类比 JS 的基本类型传参（number, string）
func doubleByValue(n int) {
	n = n * 2 // 修改的是副本
}

// 指针传递：拿到地址，可以修改原始值
// 类比 JS 传对象（引用传递的效果）
func doubleByPointer(n *int) {
	*n = *n * 2 // 解引用后修改
}

func passByValue() {
	a := 10
	doubleByValue(a)
	fmt.Println("after doubleByValue:", a) // 10，没变

	b := 10
	doubleByPointer(&b) // & 传地址
	fmt.Println("after doubleByPointer:", b) // 20，变了
}

// ============================================================
// 3. 指针接收者 vs 值接收者（现在完全理解了）
// ============================================================

type Counter struct {
	count int
}

// 值接收者：c 是副本，修改无效
func (c Counter) IncrementByValue() {
	c.count++ // 修改的是副本
}

// 指针接收者：c 是指针，修改原始 struct
func (c *Counter) Increment() {
	c.count++
}

func (c Counter) Value() int {
	return c.count
}

func receiverDemo() {
	c := Counter{}
	c.IncrementByValue()
	fmt.Println("after IncrementByValue:", c.Value()) // 0，没变

	c.Increment()
	fmt.Println("after Increment:", c.Value()) // 1，变了

	// Go 的语法糖：c.Increment() 等价于 (&c).Increment()
	// Go 自动帮你取地址，不需要手写 (&c).Increment()
}

// ============================================================
// 4. nil 指针
// 指针的零值是 nil，解引用 nil 指针会 panic
// ============================================================

func nilPointer() {
	var p *int // 零值是 nil，不是 0
	fmt.Println("p =", p) // <nil>

	// *p = 1  ← 这会 panic: nil pointer dereference

	// 使用前判断
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("p is nil, skip")
	}
}

// ============================================================
// 5. new() 创建指针
// ============================================================

func newDemo() {
	// new(T) 分配内存，返回 *T，初始值是零值
	p := new(int)
	fmt.Println("*p =", *p) // 0
	*p = 42
	fmt.Println("*p =", *p) // 42

	// 实际中更常用 & 字面量，new() 用得少
	// p2 := &int{} ← 这是错的，基本类型不能这样
	// struct 才能用字面量：p3 := &Counter{count: 0}
}

func main() {
	fmt.Println("=== 基础 ===")
	basicPointer()

	fmt.Println("\n=== 值传递 vs 指针传递 ===")
	passByValue()

	fmt.Println("\n=== 值接收者 vs 指针接收者 ===")
	receiverDemo()

	fmt.Println("\n=== nil 指针 ===")
	nilPointer()

	fmt.Println("\n=== new() ===")
	newDemo()
}
