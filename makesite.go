package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(string(data))

}
