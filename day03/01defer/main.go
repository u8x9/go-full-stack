package main

import "fmt"

// go语言中函数的return不是原子操作，在底层分为2步来执行
// 1. 返回值赋值
// 2. 真正的return返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("A, foo")
	defer fmt.Println("B, bar")
	defer fmt.Println("C, foobar")
	defer fmt.Println("D, barfoo")
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x，不是返回值
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++ // 2. defer，修改返回值：x = 6
	}()
	return 5 // 1. 返回值赋值：x = 5
	// 3. 最终返回值x=6
}
func f3() (y int) {
	x := 5
	defer func() { // 2. defer，修改x=6
		x++
	}()
	return x // 1. 返回值y=x=5
	// 3. 最终返回值y=5
}
func f4() (x int) {
	defer func(x int) { // 2. defer, 返回值x作为参数传递到匿名函数
		x++ // 操作的是参数x，而不是返回值
	}(x)
	return 5 // 1. 返回值赋值x=5
	// 3. 最终返回值x=5
}
func main() {
	deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
