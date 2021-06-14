package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "strings"
)

func main() {
	str1 := "The quick brown fox"
	str2 := "jumped over a"
	str3 := "lazy dog"
	aNum := 42
	isTrue := true

	strLen, err := fmt.Println(str1, str2, str3)
	if err == nil {
		fmt.Println("String Len", strLen)
	}

	fmt.Printf("Value of aNum: %v\n", aNum)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNum as float: %.2f\n", float64(aNum))

	fmt.Printf("Data Types: %T, %T, %T\n", str1, aNum, isTrue)

	myStr := fmt.Sprintf("Data Types var: %T, %T, %T\n", str1, aNum, isTrue)
	fmt.Print(myStr)

	//  var s string
	//  fmt.Scanln(&s)
	//  fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter Number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err == nil {
		fmt.Println("Value of Number: ",f)
	} else {
		fmt.Println(err)
	}
}
