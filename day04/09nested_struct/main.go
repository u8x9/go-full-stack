package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address // 匿名嵌套结构体
}

type company struct {
	name string
	address
}

func main() {
	p := person{
		name: "张三",
		age:  123,
		address: address{
            province: "天堂",
			city: "地狱",
		}}
	fmt.Println(p)
    fmt.Println(p.address.province)
    fmt.Println(p.province) // 先在自己结构体找这个字段，如果没有找到就在匿名嵌套结构体中找
}
