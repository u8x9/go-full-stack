package main

import "fmt"

// 递归
// 适合处理那种问题相同，问题的规模越来越小的场景

// 阶乘
func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 上台阶：n个台阶，一次可以走1步，也可以走2步，有多少种走法
func steps(n uint) uint {
	if n == 1 {
        // 如果只有一个台阶，就一种走法
		return 1
	} else if n == 2 {
        // 如果有两个台阶，就两种走法
		return 2
	}
	return steps(n-1) + steps(n-2)
}
func main() {
	fmt.Println(factorial(5))
	fmt.Println(steps(10))
}
