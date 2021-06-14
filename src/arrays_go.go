package main

import "fmt"

func main() {

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Yellow"

	fmt.Println(colors)

  var numbers = [5]int{5,4,3,2,1}
  fmt.Println(numbers)

  fmt.Println("Number of colors ", len(colors))
}
