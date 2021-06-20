package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type User struct {
	Login, Url, Admin string
	Id                int
}

func main() {

	url := "https://api.github.com/users"
	content := contentFromServer(url)
	users := usersFromJson(content)
	fmt.Println(users)
}
func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response : %T\n", resp)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	return content
}
func usersFromJson(content string) []User {
	users := make([]User, 0, 10)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var user User
	for decoder.More() {
		err := decoder.Decode(&user)
		checkError(err)
		users = append(users, user)
	}
	return users
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
