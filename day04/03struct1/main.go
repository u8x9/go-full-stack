package main

import "fmt"

type person struct {
	name   string
	age    int
	gender rune
	hobby  []string
}

func main() {
	p := person{name: "周琳", age: 32, gender: '女', hobby: []string{"吃饭", "睡觉", "打豆豆"}}
	fmt.Printf("%#v\n", p)

	// 匿名结构体
	var s struct {
		x string
		y int
	}
    s.x = "奋勇"
    s.y = 123
    fmt.Printf("type of s is %T, value is %v\n", s, s)
}
