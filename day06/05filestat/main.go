package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", fileInfo)
    fmt.Println("file size: ", fileInfo.Size())
}
