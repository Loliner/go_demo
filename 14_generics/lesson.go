package main

import "fmt"

// ============================================================
// Go 泛型（Generics）— Go 1.18+
// ============================================================
//
// JS 类比：
//   - Go 泛型 ≈ TypeScript 的泛型（T extends ...）
//   - 但 Go 是编译期静态分发，不是运行时擦除
//
// 解决的问题：
//   没有泛型时，写一个"对任意类型求和"的函数，
//   要么用 any（丢失类型安全），要么为每种类型重复写一遍。

// ============================================================
// 1. 类型参数（Type Parameter）
// ============================================================
//
// 语法：func FuncName[T 约束](参数 T) 返回值
//       类比 TS：function funcName<T extends 约束>(arg: T): T
//
// comparable 是内置约束，表示"可以用 == 比较"的类型

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// ============================================================
// 2. 自定义约束（Constraint）
// ============================================================
//
// 约束用 interface 定义，~ 表示"底层类型是..."
// 类比 TS：type Number = number | bigint

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// ============================================================
// 3. 泛型 struct
// ============================================================
//
// 类比 TS：class Stack<T> { ... }

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// ============================================================
// 4. 标准库：golang.org/x/exp/slices 之前，手写 Map/Filter
//    Go 1.21+ 标准库已经有 slices 包，但理解泛型更重要
// ============================================================

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// ============================================================
// lesson 入口
// ============================================================

func lesson() {
	// 1. 类型参数
	fmt.Println("=== 1. 类型参数 ===")
	fmt.Println(Contains([]int{1, 2, 3}, 2))       // true
	fmt.Println(Contains([]string{"a", "b"}, "c")) // false
	// 编译器能自动推断 T，不需要显式写 Contains[int](...)

	// 2. 自定义约束
	fmt.Println("\n=== 2. 自定义约束 ===")
	fmt.Println(Sum([]int{1, 2, 3}))      // 6
	fmt.Println(Sum([]float64{1.1, 2.2})) // 3.3

	// 3. 泛型 struct
	fmt.Println("\n=== 3. 泛型 struct ===")
	s := Stack[string]{}
	s.Push("a")
	s.Push("b")
	fmt.Println(s.Len()) // 2
	v, ok := s.Pop()
	fmt.Println(v, ok) // b true

	// 4. Map / Filter
	fmt.Println("\n=== 4. Map / Filter ===")
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled) // [2 4 6 8 10]

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4]

	// Map 还可以做类型转换
	strs := Map(nums, func(n int) string { return fmt.Sprintf("item-%d", n) })
	fmt.Println(strs) // [item-1 item-2 item-3 item-4 item-5]
}
