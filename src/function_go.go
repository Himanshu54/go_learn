package main

import "fmt"

func main() {
	doSomething()

	sum := addValues(10, 12)
	fmt.Println(sum)
	sum = addAllValues(12, 10, 23)
	fmt.Println(sum)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1 int, value2 int) int {
	//func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	fmt.Printf("%T\n", values)
	for i := range values {
		sum += values[i]
	}
	return sum
}
