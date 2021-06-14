package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

// methods

func (d Dog) speak() {
	fmt.Println("Bark\n")
}
func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, \n Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.speak()
}
