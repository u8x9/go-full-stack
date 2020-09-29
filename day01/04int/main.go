package main

import "fmt"

// 整型

func main(){
    var a = 10
    fmt.Printf("a 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X\n", a, a, a)

    // 八进制 
    b := 077
    fmt.Printf("b 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X\n", b, b, b)

    // 十六进制
    c := 0xff
    fmt.Printf("c 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X\n", c, c, c)
}
