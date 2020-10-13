package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行 main 函数
func main() {
	for i := 0; i < 100000; i++ {
		//go hello(i) // 开启一个单独的 goroutine 去执行 hello 函数（任务
        go func(i int) {
            fmt.Println("hello", i)
        }(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second * 3)
	// main 函数结束了，由main函数启动的所有goroutine也都结束了
}
