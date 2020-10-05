package main

import (
	"fmt"
	"strings"
)

// 闭包：是一个函数，包含了其外部作用域的变量

func adder() func(int) int {
	var x int = 100
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	fmt.Println("this is f3")
	return func() {
		f(x, y)
	}
}

// ---

func makeSuffixFn(suffix string) func(string) string {
	if !strings.HasPrefix(suffix, ".") {
		suffix = "." + suffix
	}
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += suffix
		}
		return name
	}
}

// ----

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	a := adder()
	ret := a(200)
	fmt.Println(ret)

	f1(f3(f2, 123, 456))

	jpgFn := makeSuffixFn(".jpg")
	txtFn := makeSuffixFn("txt")
	fmt.Println(jpgFn("foo"))
	fmt.Println(jpgFn("foobar.jpg"))
	fmt.Println(txtFn("bar"))

	fn1, fn2 := calc(10)
	fmt.Println(fn1(1), fn2(2)) // 11 9
	fmt.Println(fn1(3), fn2(4)) // 12 8
	fmt.Println(fn1(5), fn2(6)) // 13 7
}
