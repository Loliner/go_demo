// package 名和目录名保持一致是惯例
package mathutil

import "fmt"

// 大写开头 = exported（公开），类比 JS 的 export
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

// 小写开头 = unexported（私有），只在本包内可用，类比 JS 的不 export
func helper() {
	fmt.Println("I'm private to this package")
}
