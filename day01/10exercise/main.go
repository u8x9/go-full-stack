package main

import (
	"fmt"
	"math"
)

func main() {
    s:="Hello沙河小王子ya"
    cnLen := 0
    for _,c := range s {
        if c <= math.MaxUint8 {
            continue
        }
        cnLen++
    }
    fmt.Printf("%q中有%d个中文\n",s, cnLen)
}
