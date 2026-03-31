package main

import "fmt"

// --- 1. 命名返回值 (Named Return Values) ---
// 普通写法：返回值只有类型
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// 命名返回值写法：给返回值起名，相当于在函数顶部声明了变量
func minMaxNamed(nums []int) (min, max int) {
	min, max = nums[0], nums[0] // 注意：这里是 = 不是 :=，变量已经声明了
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // 裸 return：自动返回 min 和 max 的当前值
}

// 命名返回值的实际用途：配合 defer 修改返回值
func riskyOp() (result int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("riskyOp failed: %w", err) // %w 包装 error，类比 JS 的 new Error('...', { cause: e })
		}
	}()
	result = 42
	err = fmt.Errorf("something went wrong")
	return
}

// --- 2. Variadic 函数（可变参数）---
// 类比 JS 的 ...args 或 Python 的 *args
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 混合普通参数和可变参数
func logWithPrefix(prefix string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("[%s] %s\n", prefix, msg)
	}
}

func lesson() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

	fmt.Println("=== Named Return Values ===")
	min1, max1 := minMax(nums)
	fmt.Println("普通写法:", min1, max1)

	min2, max2 := minMaxNamed(nums)
	fmt.Println("命名返回值:", min2, max2)

	result, err := riskyOp()
	fmt.Println("riskyOp result:", result, "err:", err)

	fmt.Println("\n=== Variadic Functions ===")
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum():", sum())

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("sum(arr...):", sum(arr...))

	logWithPrefix("INFO", "服务启动", "监听 :8080")
	logWithPrefix("ERROR", "连接失败")
}
