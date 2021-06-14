package main

import (
	"fmt"
	"sort"
)

func main() {

  // array
	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
  // slices
	var numbers2 = make([]int, 5, 5)
	fmt.Println(numbers2)

	var numbers3 = []int{1, 3, 2, 5, 4}
	fmt.Println(numbers3)

  numbers2 = append(numbers2,12)
  fmt.Println(cap(numbers))
  fmt.Println(cap(numbers2))
  fmt.Println(cap(numbers3))

  sort.Ints(numbers3)
  fmt.Println(numbers3)
}
