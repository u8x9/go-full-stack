package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]

	// 删除索引为1的元素
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(a1)
	fmt.Println(s1)
}
