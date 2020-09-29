package main

import "fmt"

// 整型

func main(){
    var a = 10
    fmt.Printf("a 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X, 二进制表示为：%b\n", a, a, a, a)

    // 八进制 
    b := 077
    fmt.Printf("b 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X, 二进制表示为：%b\n", b, b, b, b)

    // 十六进制
    c := 0xff
    fmt.Printf("c 的十进制表示为：%d, 八进制表示为：%o，十六进制表示为：%X，二进制表示为：%b\n", c, c, c, c)

    // 变量的类型
    fmt.Printf("type of 'a' is %T\n", a)
    d := int8(123)
    fmt.Printf("type of 'd' is %T\n", d)
}
