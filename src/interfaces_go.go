package main

import (
	"fmt"
)

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) speak() string {
	return "Bark\n"
}

type Cat struct {
}

func (c Cat) speak() string {
	return "Meow\n"
}
func main() {
	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.speak())
	}
}
