package main

import (
	"fmt"
)

// &T{...} , &someLocalVar , new , make
func main() {

	var m map[string]int
	// m["key"] = 42
	// panic: assignment to entry in nil map
	fmt.Println(m)
	fmt.Printf("%p\n", &m)

	m = map[string]int{"key": 42}
	fmt.Println(m)
	fmt.Printf("%p\n", &m)

	m2 := make(map[string]int)
	m2["key"] = 42
	fmt.Println(m2)
	fmt.Printf("%p\n", &m2)

}
