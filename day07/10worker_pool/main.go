package main

import (
	"fmt"
	"time"
)

// worker pool(goroutine池)
// 编写代码实现一个计算随机数的每个位置数字之和的程序，要求使用 goroutine 和 channel 构建生产者和消费者模式，
// 可以指定启动的 goroutine 数量 —— worker pool 模式
// 在实际应用中，通常使用 workerpool 模式，控制 goroutine 的数量
// 防止 goroutine 泄漏和暴涨

// jobs: 只读channel
// results：只写channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker: %d start job: %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker: %d end job: %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个 goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5 个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		if i, ok := <-results; ok {
			fmt.Println("get result:", i)
		}
	}
}
