package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10) // make 之后，已经有5个值为零值的元素了
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

    // 排序
    var a1 = [...]int{3,7,8,9,1}
    sort.Ints(a1[:])
    fmt.Println(a1)
}
