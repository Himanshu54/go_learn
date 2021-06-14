package main

import (
	"fmt"
)

var isConnected bool = true

// defer
//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 
func main() {
	deferedStatements()
	doSomething()
	fmt.Println("isConnected after doSomenthing:", isConnected)

}

func deferedStatements() {
	defer fmt.Println("Statement 1")
	i := 1
	defer fmt.Println("I1: ", i)
	i++
	fmt.Println("I2: ", i)
	fmt.Println("Statement 2")
}

func doSomething() {
	defer disconnect()
	fmt.Println("isConnected after defer disconnect :", isConnected)
}

func disconnect() {
	isConnected = false
}
