package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.github.com/users/defunkt"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response : %T", resp)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	fmt.Print(content)
}
