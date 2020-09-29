package main

import "fmt"

func main() {
    var n = 28734

    fmt.Printf("类型：%T\n", n) // 类型
    fmt.Printf("变量的值：%v\n", n) // 变量的值 --> 任意变量
    fmt.Printf("二进制：%b\n", n) // 二进制
    fmt.Printf("十进制：%d\n", n) // 十进制
    fmt.Printf("八进制：%o\n", n) // 八进制
    fmt.Printf("十六进制，小写：%x; 大写：%X\n", n, n) // 十六进制

    var s = "Hello 沙河！"
    fmt.Printf("字符串：%s\n", s) // 字符串
    fmt.Printf("字符串：%v\n", s) // 字符串
    fmt.Printf("字符串带引号：%q\n", s) // 字符串
    fmt.Printf("字符串带引号：%#v\n", s) // 字符串
}
