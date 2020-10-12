package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeAndWriteString() {
	file, err := os.OpenFile("foo.bar", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write([]byte("Hello, 世界\n"))
	file.WriteString("你好，World\n")
}

func bufioNewWriter() {
	file, err := os.OpenFile("foo.bar", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write([]byte("春眼不觉晓\n")) // 写到缓存
	writer.WriteString("处处闻啼鸟\n")
	writer.Flush() // 刷新缓存，写入到文件中
}
func ioutilWrite() {
	buf := []byte("夜来风雨声\n花落知多少\n")
	if err := ioutil.WriteFile("foo.bar", buf, 0644); err != nil {
		fmt.Println(err)
	}
}
func main() {
	//writeAndWriteString()
	//bufioNewWriter()
    ioutilWrite()
}
