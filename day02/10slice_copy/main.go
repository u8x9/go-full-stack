package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5,6}
	s2 := make([]int, 6)
	copy(s2, s1)
    s1[2] = 999
	fmt.Println(s1, s2)

    // 将 s1 中索引为1的元素删除
    s1 = append(s1[:1], s1[2:]...)
    fmt.Println(s1)
}
