package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, 世界")

	// Array comparision needs order.
	var a = [...]int{1, 2, 3}
	var b = [...]int{3, 2, 1}
	fmt.Println(a == b)

	b = [...]int{1, 2, 3}
	fmt.Println(a == b)

	// Slice are also ordered.
	var x = []int{1, 2, 3}
	var y = []int{3, 2, 1}
	fmt.Println(slices.Equal(x, y))

	y = append(x, 3, 2, 1)
	fmt.Println(x)
	fmt.Println(y)
}
