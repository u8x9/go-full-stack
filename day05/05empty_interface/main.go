package main

import (
	"fmt"
	"strings"
)

func show(x interface{}) {
	fmt.Printf("%T, %v\n", x, x)
}

func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "张三"
	m1["age"] = 123
	m1["gender"] = true
	m1["hobby"] = [...]string{"吃饭", "睡觉", "打豆豆"}

	show(m1)
	show(false)
	show(nil)

	// 类型断言

	var xx interface{} = "Hello"
	if v, ok := xx.(string); ok {
		fmt.Println(strings.ToUpper(v))
	}
}
