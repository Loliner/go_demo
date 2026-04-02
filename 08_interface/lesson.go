package main

import (
	"fmt"
	"math"
)

// ============================================================
// Go Interface 核心概念
// ============================================================
//
// JS 类比：interface 类似于 TypeScript 的 interface，但有一个关键区别：
//   - TS: 需要显式声明 "implements SomeInterface"
//   - Go: 隐式实现，只要你有对应的方法，就自动满足 interface
//
// 这叫 "duck typing"（鸭子类型）：像鸭子一样走路、像鸭子一样叫，那就是鸭子。

// ============================================================
// 1. 定义 interface
// ============================================================

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 实现了 Shape interface（隐式，不需要任何声明）
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 也实现了 Shape interface
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 接受 interface 类型的函数 —— 可以传入任何实现了 Shape 的类型
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ============================================================
// 2. interface 可以存储任何实现了它的类型
// ============================================================

func demoPolymorphism() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 1},
	}

	for _, s := range shapes {
		printShapeInfo(s)
	}
}

// ============================================================
// 3. 空 interface —— any（Go 1.18 之前写 interface{}）
//
// 类比 JS 的 any / TypeScript 的 unknown
// 可以存储任意类型的值
// ============================================================

func printAnything(v any) {
	fmt.Printf("value: %v, type: %T\n", v, v)
}

func demoAny() {
	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything([]int{1, 2, 3})
}

// ============================================================
// 4. Type Assertion —— 从 interface 取回具体类型
//
// 类比 TypeScript 的类型断言 as，但 Go 在运行时检查
// ============================================================

func demoTypeAssertion() {
	var s Shape = Circle{Radius: 3}

	// 断言为 Circle（安全写法，带 ok）
	if c, ok := s.(Circle); ok {
		fmt.Printf("It's a Circle with radius %.1f\n", c.Radius)
	}

	// 断言失败不会 panic（ok 为 false）
	if _, ok := s.(Rectangle); !ok {
		fmt.Println("It's not a Rectangle")
	}

	// 危险写法（不带 ok），断言失败会 panic
	// c := s.(Rectangle) // 别这样写
}

// ============================================================
// 5. Type Switch —— 多类型分支，比一堆 if.(type) 优雅
//
// 类比 JS 的 switch(typeof x) 但更强大，能匹配具体类型
// ============================================================

func describe(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case bool:
		return fmt.Sprintf("bool: %v", v)
	case Circle:
		return fmt.Sprintf("Circle with radius %.1f", v.Radius)
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

func demoTypeSwitch() {
	fmt.Println(describe(42))
	fmt.Println(describe("hello"))
	fmt.Println(describe(Circle{Radius: 2}))
	fmt.Println(describe(3.14))
}

// ============================================================
// 6. 组合 interface（interface embedding）
//
// 类比 TypeScript 的 interface 继承（extends）
// ============================================================

type Stringer interface {
	String() string
}

type Saver interface {
	Save() error
}

// 组合两个 interface
type StringerSaver interface {
	Stringer
	Saver
}

// ============================================================
// 7. 标准库里最重要的 interface：fmt.Stringer
//
// 只要实现了 String() string，fmt.Println 就会自动调用它
// 类比 JS 的 toString()
// ============================================================

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func demoStringer() {
	p := Point{3, 4}
	fmt.Println(p)        // 自动调用 String()，输出 (3, 4)
	fmt.Printf("%v\n", p) // 同上
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	fmt.Println("=== 1. Polymorphism ===")
	demoPolymorphism()

	fmt.Println("\n=== 2. any (empty interface) ===")
	demoAny()

	fmt.Println("\n=== 3. Type Assertion ===")
	demoTypeAssertion()

	fmt.Println("\n=== 4. Type Switch ===")
	demoTypeSwitch()

	fmt.Println("\n=== 5. fmt.Stringer ===")
	demoStringer()
}
