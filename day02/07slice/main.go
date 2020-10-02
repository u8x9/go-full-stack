package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"张三", "李四", "王五"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)

	// 长度和容量
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1))

	// 基于数组定义切片
	a := [5]int{1, 2, 3, 4, 5}
    s := a[1:4] // 左闭右开
	fmt.Printf("value of s is %v, type of s is %T\n", s, s)
	fmt.Printf("len of s is %d, cap of s is %d\n", len(s), cap(s))
	s = a[:]
	fmt.Println(s)
	s = a[:3]
	fmt.Println(s)
	s = a[3:]
	fmt.Println(s)
}
