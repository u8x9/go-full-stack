package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	goroutine 和 channel 练习

	使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序
	1. 开启一个 goroutine 循环生成 int64 类型的随机数，发送到 jobChan
	2. 开启24个 goroutine 从 jobChan 中取出随机数，计算各位数的和，将结果发送到 resultChan
	3. 主 goroutine 从 resultChan 取出结果并打印到终端
*/

func gen(jobChan chan<- int64) {
	for {
		jobChan <- rand.Int63()
        time.Sleep(time.Microsecond * 300)
	}
}
func calc(jobChan <-chan int64, resultChan chan<- int64) {
	for i := range jobChan {
		resultChan <- sum(i)
        time.Sleep(time.Microsecond * 300)
	}
}
// 计算各位数的和
func sum(i int64) int64 {
    n := i
    s := int64(0)
    for {
        if n <= 0 {
            break
        }
        j := n % 10
        n /= 10
        s += j
    }
    return s
}
func main() {
	rand.Seed(time.Now().UnixNano())

	jobChan := make(chan int64, 100)
	resultChan := make(chan int64, 100)

	go gen(jobChan)
	for i := 0; i < 24; i++ {
		go calc(jobChan, resultChan)
	}

	for i := range resultChan {
		fmt.Println("result:", i)
	}

}
