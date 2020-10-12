package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println("打开文件失败：", err.Error())
		return
	}
	var buffer = make([]byte, 128)
	file.Seek(0, os.SEEK_SET)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件失败：", err.Error())
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	defer file.Close()
}
