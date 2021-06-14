package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	fmt.Println(sum)

	colors := []string{"Red", "Green", "Blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	for i := range colors {
		fmt.Println(i, ": ", colors[i])
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 500 {
			break
		}
		if sum > 200 {
			goto endofprogram
		}
	}
	fmt.Println(sum)

endofprogram:
	fmt.Println("End of Program")
}
