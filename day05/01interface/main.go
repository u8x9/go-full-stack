package main

import "fmt"

type speaker interface {
	speak()
}

type cat struct{}
type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵～")
}
func (d dog) speak() {
	fmt.Println("汪汪汪～")
}

func beat(s speaker) {
	s.speak()
}

func main() {
	c := cat{}
	d := dog{}
	beat(c)
	beat(d)
}
