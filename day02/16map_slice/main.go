package main

import "fmt"

// map 和 slice 组合

func main(){
    // 元素类型为map的切片
    s1 := make([]map[string]int, 10, 10)
    s1[0] = make(map[string]int, 1)
    s1[0]["foo"] = 99
    fmt.Println(s1)

    // 值为切片的map
    m1 := make(map[int][]int, 1)
    m1[0] = []int{1,2,3}
    fmt.Println(m1)
}
