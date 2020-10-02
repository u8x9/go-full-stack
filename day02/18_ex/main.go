package main

import "fmt"

func main(){
    // 写一个程序，统计一个字符串中，每个字符出现的次数
    s := "坡上立着一只鹅，坡下就是一条河。宽宽的河，肥肥的鹅，鹅要过河，河要渡鹅不知是鹅过河，还是河渡鹅Peter Piper picked a peck of pickled peppers. A peck of pickled peppers Peter Piper picked."
    m := make(map[rune]int)
    for _,c := range s {
        m[c]++
    }
    for k, v := range m {
        fmt.Printf("%c: %02d\n", k, v)
    }
}
