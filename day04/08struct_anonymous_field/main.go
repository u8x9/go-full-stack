package main

import "fmt"

// 结构体的匿名字段

type person struct {
    string
    int
}

func main(){
    p := person{"张三", 123}
    fmt.Printf("%#v\n", p)
    fmt.Println(p.string, p.int)
}
