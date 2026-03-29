package main

import "fmt"

func stats(nums ...int) (sum int, avg float64, err error) {
	if len(nums) == 0 {
		err = fmt.Errorf("no numbers provided")
		return
	}
	// 你的实现
	for _, n := range nums {
		sum += n
	}
	avg = float64(sum) / float64(len(nums))
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum, avg, err := stats(nums...)
	fmt.Printf("Sum: %d, Average: %.2f, Error: %v\n", sum, avg, err)
}
