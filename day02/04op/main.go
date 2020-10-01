package main

import (
	"fmt"
)

// 运算符

func main() {
	flag := 4

	// 算术运算符
	if flag == 1 {
		var (
			a = 5
			b = 2
		)
		fmt.Println(a + b)
		fmt.Println(a - b)
		fmt.Println(a * b)
		fmt.Println(a / b)
		fmt.Println(a % b)
		a++ // 单独的语句，不能作为右值
		b-- // 单独的语句，不能作为右值
		fmt.Println(a, b)
	}

	// 关系运算符
	// 只有相同数据类型的才能进行比较
	if flag == 2 {
		var (
			a = 5
			b = 2
		)
		fmt.Println(a == b)
		fmt.Println(a != b)
		fmt.Println(a > b)
		fmt.Println(a >= b)
		fmt.Println(a < b)
		fmt.Println(a <= b)
	}

	// 逻辑运算符
	if flag == 3 {
		age := 22
		if age > 18 && age < 60 {
			fmt.Println("苦逼的赚钱")
		} else {
			fmt.Println("愉快的玩耍")
		}

		fmt.Println(age > 18 || age < 60)
		fmt.Println(!(age < 18))
	}

	// 位运算符
	if flag == 4 {
		const (
			five = 5 // 101
			two  = 2 // 010
		)
        result := five & two
        fmt.Printf("%b, %d\n", result, result)
        result = five | two
        fmt.Printf("%b, %d\n", result, result)
        result = five ^ two
        fmt.Printf("%b, %d\n", result, result)
        result = five << two
        fmt.Printf("%b, %d\n", result, result)
        result = five >>  two
        fmt.Printf("%b, %d\n", result, result)
	}
}
