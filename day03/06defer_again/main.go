package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Printf("idx=%q a=%d b=%d ret=%d\n", index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	/*
	   idx="10" a=1 b=2 ret=3
	   idx="20" a=0 b=2 ret=2
	   idx="2" a=0 b=2 ret=2
	   idx="1" a=1 b=3 ret=4
	*/
}
