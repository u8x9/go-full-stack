package main

import "fmt"

func main() {
	// 1. &: 取地址
	n := 18
	p := &n
	fmt.Printf("type of p is %T, and value is %v\n", p, p)

	// 2. *：根据地址取值
    m := *p
    fmt.Printf("type of m is %T, and value is %v\n", m, m)

    *p = 999
    fmt.Println(n, m)

    // new 函数：申请一块内存
    var a1 *int // nil
    fmt.Println(a1)
    //*a1 = 123 // panic: 空指针
    var a2 = new(int)
    fmt.Println(a2)
    *a2 = 123

    // make 也是用于内存分配，区别于new，它只用于slice, map以及chan的内存创建
    // 而且它返回的类型就是这三个类型本身，而不是它们的指针
    // 因为这三个类型本身就是引用类型，所以就没必要再返回它们的指针了
}
