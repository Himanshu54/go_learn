package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello World!"

	file, err := os.Create("./fileFromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("NO error", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./fileBytes.txt", bytes, 0644)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
