package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buf, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(buf))
}
