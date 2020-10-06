package main

import "fmt"

// 结构体是值类型

type Person struct {
	Name string
	Age  uint8
}

// go语言中，函数函数永远是值传递
func f(p *Person) {
	p.Age = 123
}
func main() {
	p := Person{}
	p.Name = "张三"
	p.Age = 10
	f(&p)
	fmt.Println(p)

	p1 := new(Person)
	p1.Name = "李四"
	f(p1)
	fmt.Println(p1)

	// key-value 初始化
	p2 := &Person{
		Name: "王五",
		Age:  23,
	}
	fmt.Println(p2)

	// 使用列表形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p3 := &Person{
		"赵六",
		32,
	}
	fmt.Println(p3)
}
