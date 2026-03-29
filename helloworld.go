package main

import "fmt"

func main() {
	// 创建 1-15 的 slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n\n", cap(numbers))

	// 切片：共享底层数组
	neededNumbers := numbers[:len(numbers)-10]
	neededNumbers2 := numbers[:len(numbers)-10]
	neededNumbers[0] = 99
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("neededNumbers: %v\n", neededNumbers)
	fmt.Printf("neededNumbers2: %v\n\n", neededNumbers2)

	// copy：独立内存
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	numbersCopy[0] = 100

	fmt.Printf("\nnumbersCopy = %v\n", numbersCopy)
	fmt.Printf("numbersCopy length = %d\n", len(numbersCopy))
	fmt.Printf("numbersCopy capacity = %d\n", cap(numbersCopy))
}
