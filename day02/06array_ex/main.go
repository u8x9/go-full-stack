package main

import "fmt"

// 数组练习题

func main() {
	// 1. 求数组[1,3,5,7,8]所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, i := range a1 {
		sum += i
	}
	fmt.Println(sum)

	// 2. 找出数组中和为指定值的两个无素的下标。比如从数组[1,3,5,7,8]找出和为8的两个元素下标分别是(0,3)和(1,2)
	target := 8
	arrLen := len(a1)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if a1[i]+a1[j] == target {
                fmt.Printf("(%d, %d)\n", i, j)
				break
			}
		}
	}
}
