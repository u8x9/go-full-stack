package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := int32(31784)
	//s:=string(i) // 错误的方法
	s := fmt.Sprintf("%d", i)
	fmt.Println(s)

	s = "9527"
	ii, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of ii: %T, value of ii: %v\n", ii, ii)

	ui, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of ui: %T, value of ui: %v\n", ui, ui)

	s = "true"
	b, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of b: %T, value of b: %v\n", b, b)

	s = "3.14"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of f: %T, value of f: %v\n", f, f)

	s = "1234"
	j, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of j: %T, value of j: %v\n", j, j)

    s = strconv.Itoa(j)
    fmt.Printf("type of s: %T, value of s: %v\n", s,s)
}
