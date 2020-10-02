package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(a, s)
	s = append(s, 4, 5, 6)
	fmt.Println(a, s)

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	s1 := []string{"天津", "日照", "海口"}
	s2 := []string{"重庆", "苏州", "西安"}
	s1 = append(s1, s2...) // ... 表示拆开
	fmt.Println(s1)
}
