package main

import "fmt"

func stats(nums ...int) (sum int, avg float64, err error) {
	if len(nums) == 0 {
		err = fmt.Errorf("no numbers provided")
		return
	}
	for _, n := range nums {
		sum += n
	}
	avg = float64(sum) / float64(len(nums))
	return
}

func exercise() {
	fmt.Println("=== Exercise: stats ===")

	nums := []int{1, 2, 3, 4, 5}
	s, avg, err := stats(nums...)
	fmt.Printf("Sum: %d, Average: %.2f, Error: %v\n", s, avg, err)

	_, _, err = stats()
	fmt.Println("empty stats error:", err)
}
