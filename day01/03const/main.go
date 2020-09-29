package main

import "fmt"

const PI = 3.1415926

// 批量声明
const (
	statusOK = 200
	notFound = 404
)

const (
	n1 = 100
	n2 // 100，如果没有指定值，和上面的值一样
	n3 // 100
)

// iota 常量计数器
// 在 const 关键字出现时，被重置为0
// const 中每新增一行常量声明，将使用 iota 计数一次
const (
	m1 = iota // 0
	m2        // 1
	m3        // 2
)

const (
	a1 = iota // 0
	a2        // 1
	_         // 2
	a3        // 3
)

const (
	b1 = iota // 0
	b2 = 100  // 100, iota = 1
	b3 = iota // 2
	b4        // 3
)

const (
	// !!! 每新增一行常量声明，iota 的值加1 !!!
	c1, c2 = iota + 1, iota + 2 // 本行 iota == 0
	c3, c4 = iota + 1, iota + 2 // 本行 iota == 1
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	flag := 5
	if flag == 0 {
		fmt.Println("n1: ", n1)
		fmt.Println("n2: ", n2)
		fmt.Println("n3: ", n3)
	}

	if flag == 1 {
		fmt.Println("m1: ", m1)
		fmt.Println("m2: ", m2)
		fmt.Println("m3: ", m3)
	}

	if flag == 2 {
		fmt.Println("a1: ", a1)
		fmt.Println("a2: ", a2)
		fmt.Println("a3: ", a3)
	}

	if flag == 3 {
		fmt.Println("b1: ", b1)
		fmt.Println("b2: ", b2)
		fmt.Println("b3: ", b3)
		fmt.Println("b4: ", b4)
	}

	if flag == 4 {
		fmt.Println("c1: ", c1)
		fmt.Println("c2: ", c2)
		fmt.Println("c3: ", c3)
		fmt.Println("c4: ", c4)
	}

	if flag == 5 {
		fmt.Println("KB: ", KB)
		fmt.Println("MB: ", MB)
		fmt.Println("GB: ", GB)
		fmt.Println("TB: ", TB)
		fmt.Println("PB: ", PB)
	}
}

