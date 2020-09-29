package main

import "fmt"

// go语言中推荐使用驼峰式命名
//var student_name string
var studengName string

//var StudentName string

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 16
	isOk = true

	fmt.Println(isOk)
	fmt.Printf("name: %s\n", name)
	fmt.Println(age)

	// 声明变量同时赋值
	var s1 string = "foo"
	fmt.Println(s1)

	// 类型推导
	var s2 = "bar"
	fmt.Println(s2)

	// 简短变量声明（不能用于全局变量）
	s3 := "foobar"
	fmt.Println(s3)
}
