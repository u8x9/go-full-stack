package main

import "fmt"

func fa() {
	fmt.Println("a")
}

func fb() {
	// 1. recover() 必须搭配defer使用
	// 2. recover() 所在的 defer 一定要在可能引用painc的语句之前定义
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到错误：", err)
		}
	}()
	panic("出现致命错误!") // 程序崩溃退出
	fmt.Println("b")
}
func fc() {
	fmt.Println("c")
}

func main() {
	fa()
	fb()
	fc()
}
