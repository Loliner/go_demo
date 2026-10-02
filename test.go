package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:3]
	c := a[3:]
	b = append(b, 100)
	c = append(c, 200)
	c[0] = -1
	fmt.Println(a)
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}

