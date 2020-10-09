package main

import "fmt"

// 结构体模拟实现继承

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%q在移动\n", a.name)
}

type dog struct {
	feet uint8
    name string
	animal
	//*animal
}

func (d *dog) wang() {
	fmt.Printf("%q汪汪叫\n", d.name)
}


func main(){
    d := dog {
        feet: 4,
        name: "橘子汁",
        animal: animal {
        //animal: &animal {
            name: "浪花",
        },
    }
    fmt.Printf("%#v\n", d)
    d.wang()
    d.move()
}
