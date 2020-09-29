package main

import (
	"fmt"
)

// go语言为了处理非ASCII字符，定义了 rune 类型

func main(){
    s := "Hello沙河"
    // len() 求的是byte字节的数量
    n := len(s)
    fmt.Println(n)

    // byte

//    for i:= 0; i < len(s); i++ {
//        //fmt.Println(s[i])
//        fmt.Printf("%c\n", s[i])
//    }


    // rune 
    for _, c := range s {
        //fmt.Println(c)
        fmt.Printf("%c\n", c)
    }

    // 字符串修改
    // 需要转成 []byte 或 []rune 再进行修改
    ss := []rune(s) // 把字符串强转为rune切片
    ss[5] = '天'
    fmt.Println(string(ss)) // 把[]rune 转成字符串
}
