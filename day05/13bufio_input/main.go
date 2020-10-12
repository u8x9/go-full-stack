package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var input string
	fmt.Print("请输入内容：")
	fmt.Scanln(&input) //如果有空格 会截断
	fmt.Print("你输入的内容是：", input)
}

func useBufio() {
	var input string
	fmt.Print("请输入内容：")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("你输入的内容是：", input)
}
func bufioByStdout() {
	w := bufio.NewWriter(os.Stdout)
	_, err := w.WriteString("你好，世界\n")
	if err != nil {
		return
	}
	w.Flush()
}
func main() {
	//useScan()
	//useBufio()
	bufioByStdout()
}
