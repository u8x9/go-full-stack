package main

import "fmt"

// sum：计算和
func sum(x, y int) int {
	return x + y
}
func sub(x, y int) (ret int) {
	ret = x - y
	return
}
func foo() (int, string) {
	return 123, "foobar"
}
func bar(x int, y ...int) int {
	for v := range y {
		x += v
	}
	return x
}
func main() {
	fmt.Println(sum(123, 456))
	fmt.Println(sub(123, 456))
	m, n := foo()
	fmt.Println(m, n)
	fmt.Println(bar(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
