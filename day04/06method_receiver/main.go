package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

func (d *dog) wang() {
	fmt.Println(d.name, "汪汪汪～")
}

// -----

type u8 uint8

func (u *u8) plus(other u8) {
	*u = *u + other
}

func newU8(u int) u8 {
	return u8(uint8(u))
}

func main() {
	d := newDog("旺财")
	d.wang()

	var u u8 = 12
	var w = newU8(34)
	u.plus(w)
	fmt.Println(u)
}
