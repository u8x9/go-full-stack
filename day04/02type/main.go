package main

import "fmt"

type u8 uint8     // 自定义类型
type u32 = uint32 // 类型别名

func main() {
	var n u8 = 100
	var m u32 = 777
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)
}
