package main

import "fmt"

func main(){
    age := 19
    //if age > 18 {
        //fmt.Println("欢迎光临")
    //} else {
        //fmt.Println("吃多两年饭再说")
    //}
    if age > 35  {
        fmt.Println("五好老人")
    } else if age > 18 {
        fmt.Println("四好青年")
    } else {
        fmt.Println("三好学生")
    }

    // 特殊写法
    // 【作用域】此处age1只在if语句中有效
    if age1:=19; age1 > 18 {
        fmt.Println("欢迎光临")
    } else {
        fmt.Println("吃多两年饭再说")
    }
}
