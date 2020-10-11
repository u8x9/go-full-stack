package main

import "fmt"

type canFly interface {
	fly()
}
type canSwim interface {
	swim()
}

type flyingFish struct{}

func (ff flyingFish) fly() {
	fmt.Println("飞鱼在飞")
}

func (ff flyingFish) swim() {
	fmt.Println("飞鱼在游")
}

func main() {
	var ff = flyingFish{}
	var f canFly = ff
	var s canSwim = ff

	f.fly()
	s.swim()
}
