package main

import (
	"fmt"
	"math"
)

func main(){
    fmt.Println("max value of float32 is:", math.MaxFloat32)

    f1 := 123.456 // 默认为 float64
    f2 := float32(789.10) // float32

    fmt.Printf("type of f1 is %T, type of f2 is %T\n", f1, f2)

     //f1 = f2 // float32 和 float64 之间不能直接赋值

}
