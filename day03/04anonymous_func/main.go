package main

import "fmt"

// 匿名函数

func main() {
	// 函数内部无法声明具名函数
	// 但可以声明匿名函数

	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Println(f1(123, 456))

	// 如果只是调用一次，还可以简写成立即执行函数
	func(name string) {
		fmt.Println("Hello,", name)
	}("张三")
}
