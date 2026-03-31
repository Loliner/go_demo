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

	*p = 100
	fmt.Println("x =", x) // 100，x 被修改了
}

// ============================================================
// 2. 为什么需要指针：函数传参默认是值拷贝
// ============================================================

// 值传递：拿到的是副本，修改不影响原始值
func doubleByValue(n int) {
	n = n * 2
}

// 指针传递：拿到地址，可以修改原始值
func doubleByPointer(n *int) {
	*n = *n * 2
}

func passByValue() {
	a := 10
	doubleByValue(a)
	fmt.Println("after doubleByValue:", a) // 10，没变

	b := 10
	doubleByPointer(&b)
	fmt.Println("after doubleByPointer:", b) // 20，变了
}

// ============================================================
// 3. 指针接收者 vs 值接收者
// ============================================================

type Counter struct {
	count int
}

// 值接收者：c 是副本，修改无效
func (c Counter) IncrementByValue() {
	c.count++
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

	// Go 语法糖：c.Increment() 等价于 (&c).Increment()
}

// ============================================================
// 4. nil 指针
// ============================================================

func nilPointer() {
	var p *int // 零值是 nil，不是 0
	fmt.Println("p =", p)

	// *p = 1  ← 这会 panic: nil pointer dereference

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
	// struct 才能用字面量：p2 := &Counter{count: 0}
}

func lesson() {
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
