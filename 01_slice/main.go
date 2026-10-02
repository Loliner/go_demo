package main

import "fmt"

func main() {
	// 创建 1-15 的 slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)         // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	fmt.Printf("length = %d\n", len(numbers))     // 15
	fmt.Printf("capacity = %d\n\n", cap(numbers)) // 15

	// 切片：共享底层数组
	neededNumbers := numbers[:len(numbers)-10]
	neededNumbers2 := numbers[:len(numbers)-10]
	neededNumbers[0] = 99
	fmt.Printf("numbers: %v\n", numbers)                 // [99 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	fmt.Printf("neededNumbers: %v\n", neededNumbers)     // [99 2 3 4 5]
	fmt.Printf("neededNumbers2: %v\n\n", neededNumbers2) // [99 2 3 4 5]

	// copy：独立内存
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	numbersCopy[0] = 100

	fmt.Printf("\nnumbersCopy = %v\n", numbersCopy)             // [100 2 3 4 5]
	fmt.Printf("numbersCopy length = %d\n", len(numbersCopy))   // 5
	fmt.Printf("numbersCopy capacity = %d\n", cap(numbersCopy)) // 5
}

// append 在有cap余量的时候直接修改，没余量才copy并分配新数组
// s[2:4:4] 格式是 s[起点:终点:cap上限],取下标 2 到 3 的元素(不含 4),并把 cap 限制到下标 4 为止
// 分片时，cap 是从起点一直算到底层数组末尾:cap = 原 cap - 起点下标
