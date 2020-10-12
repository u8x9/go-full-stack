package main

import (
	"fmt"

	"github.com/u8x9/go-full-stack/day05/calc"
)

func init(){
    fmt.Println("init from main")
}

func main() {
	sum := calc.Add(123, 456)
	fmt.Println(sum)
}
