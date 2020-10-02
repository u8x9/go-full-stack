package main

import "fmt"

type calcFun func(int, int) int

//func calc(f func(int, int) int, x, y int) int {
func calc(f calcFun, x, y int) int {
	return f(x, y)
}
func sum(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func foo(x func(int) int) func(int, int) int {
	return func(a, b int) int {
		return x(a) + b
	}
}
func bar(x int) int {
	return x * x
}

func main() {
	fmt.Println(calc(sum, 123, 456))
	fmt.Println(calc(sub, 123, 456))
	r := foo(bar)(3, 10)
	fmt.Println(r)
}
