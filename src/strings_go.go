package main

import (
"fmt"
"strings"
)

func main(){
str := "This strings";
fmt.Printf("str : %v, %T\n", str,str)

fmt.Println(strings.ToUpper(str))
fmt.Println(strings.Title(str))

lvalue := "hello"
rvalue := "HELLO"

fmt.Printf("%v equals %v: %v \n",lvalue ,rvalue ,(lvalue == rvalue))
fmt.Printf("%v i equals %v: %v \n",lvalue ,rvalue ,strings.EqualFold(lvalue,rvalue))

fmt.Printf("%v contains 'his' : %v \n",str,strings.Contains(str,"his"))
}
