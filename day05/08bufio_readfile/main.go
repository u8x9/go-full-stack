package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio 按行读取

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 是char
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}
