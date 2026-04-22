package main

import "fmt"

// ============================================================
// 练习：实现泛型工具函数
// ============================================================
//
// 1. 实现 Keys[K, V]：
//    接收 map[K]V，返回所有 key 组成的 []K
//    约束：K 需要 comparable
//
//    示例：
//      Keys(map[string]int{"a": 1, "b": 2}) → ["a", "b"]（顺序不定）
//
// 2. 实现 Reduce[T, U]：
//    接收 []T、初始值 U、累加函数 func(U, T) U，返回 U
//    类比 JS 的 Array.prototype.reduce
//
//    示例：
//      Reduce([]int{1,2,3}, 0, func(acc, n int) int { return acc + n }) → 6
//      Reduce([]string{"a","b","c"}, "", func(acc, s string) string { return acc + s }) → "abc"
//
// 3. 实现泛型 Pair[A, B] struct：
//    包含 First A 和 Second B 两个字段
//    实现 Swap() 方法，返回 Pair[B, A]（交换两个字段）
//
//    示例：
//      p := Pair[string, int]{First: "hello", Second: 42}
//      p.Swap() → Pair[int, string]{First: 42, Second: "hello"}

// TODO: 在这里写你的代码

func Keys[K comparable, V any](m map[K]V) []K {
	// panic("not implemented")
	var result []K
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Reduce[T, U any](slice []T, init U, fn func(U, T) U) U {
	// panic("not implemented")
	result := init
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

type Pair[A, B any] struct {
	First  A
	Second B
}

func (p Pair[A, B]) Swap() Pair[B, A] {
	// panic("not implemented")
	result := Pair[B, A]{
		First:  p.Second,
		Second: p.First,
	}
	return result
}

func exercise() {
	fmt.Println("=== Exercise: 泛型工具函数 ===")

	// 测试 1: Keys
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Keys:", Keys(m))

	// 测试 2: Reduce
	sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum) // 15

	concat := Reduce([]string{"Go", " ", "is", " ", "fun"}, "", func(acc, s string) string { return acc + s })
	fmt.Println("Concat:", concat) // "Go is fun"

	// 测试 3: Pair
	p := Pair[string, int]{First: "hello", Second: 42}
	fmt.Println("Pair:", p)
	fmt.Println("Swapped:", p.Swap())
}
