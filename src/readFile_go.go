package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "./fileBytes.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes Read: ", content)

	result := string(content)
	fmt.Println("String : ", result)
}
