package main

import (
	"fmt"
	"sync"
)

// channel 是引用类型，需要手动开辟存储空间

var ch chan int // nil
var wg sync.WaitGroup

func main() {
	// 初始化
	// make(chan 元素类型[, 缓冲大小])
	//ch = make(chan int, 1) //带缓冲区的channel
	ch = make(chan int) // 不带缓冲区的channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := <-ch // 接收数据
		fmt.Println("i:", i)
	}()
	ch <- 12 // 发送数据
	wg.Wait()
	close(ch) // 关闭channel
}
