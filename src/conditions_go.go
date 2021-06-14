package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// if
	// var x float64 = 42
	var result string

	if x := 42; x < 0 {
		result = "Less than 0"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println("Result: ", result)

	// switch
	rand.Seed(time.Now().Unix())
	//	dow := rand.Intn(6) + 1
	//	fmt.Println("Day", dow)

	result = ""
	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Sunday"
	case 7:
		result = "Saturday"
	default:
		result = "Weekday"
	}
	fmt.Println("Day", result)

	switch x := -42; {
	case x < 0:
		fmt.Println("less than 0")
		fallthrough
	case x == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Greater than 0")
	}
}
