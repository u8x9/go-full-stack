package main

import "fmt"

// 实现多个接口和接口嵌套

type mover interface {
	move()
}

type eater interface {
	eat(food string)
}

type cat struct {
	name string
}

// 一个结构体实现多个接口

func (c *cat) move() {
	fmt.Printf("%q在走猫步\n", c.name)
}

func (c *cat) eat(food string) {
	fmt.Printf("%q在吃%q，喵喵喵～\n", c.name, food)
}

type fish struct{}

// 一个接口被多个结构体实现

func (f *fish) eat(food string) {
	fmt.Println("一条小鱼很欢快的在吃", food)
}

// 接口嵌套

type animal interface {
	mover
	eater
}

func main() {
	var m mover
	var e eater
	c := &cat{
		name: "tom",
	}
	m = c
	e = c
	m.move()
	e.eat("黄花鱼")

	e = &fish{}
	e.eat("水草")

	var a animal = c
	a.move()
	a.eat("黄花鱼")
}
