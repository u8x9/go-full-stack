package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// -- 值接收者
func (c cat) move() {
	fmt.Printf("%q晃动%d只脚在走猫步\n", c.name, c.feet)
}
func (c cat) eat(food string) {
	fmt.Printf("%q吃了%q, 喵喵喵～\n", c.name, food)
}

type dog struct{}

// --- 指针接收者
func (d *dog) move() {
	fmt.Println("手牵手一步两步三步往前走")
}
func (d *dog) eat(food string) {
	fmt.Println("直勾勾的对着", food, "汪汪汪～")
}

func main() {
    var a animal
	c1:=cat{"tom",4}
    a = c1
    a.move()
    a.eat("鱼仔")

    c2:=&cat{"呀呀",4}
    a = c2
    a.move()
    a.eat("鱼仔")

	//d1:=dog{}
    //a = d1 // dog does not implement animal (eat method has pointer receiver)
    //a.move()
    //a.eat("大骨头")

	d2 := &dog{}
    a = d2
    a.move()
    a.eat("小骨头")
}
