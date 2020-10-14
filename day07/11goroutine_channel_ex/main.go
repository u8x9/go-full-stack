package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	【讲师实现] goroutine 和 channel 练习

	使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序
	1. 开启一个 goroutine 循环生成 int64 类型的随机数，发送到 jobChan
	2. 开启24个 goroutine 从 jobChan 中取出随机数，计算各位数的和，将结果发送到 resultChan
	3. 主 goroutine 从 resultChan 取出结果并打印到终端
*/

var wg sync.WaitGroup

// Job 随机数
type Job struct {
	Value int64
}

// NewJob Job构造函数
func NewJob(x int64) *Job {
	return &Job{
		Value: x,
	}
}

// Result 结果
type Result struct {
	*Job
	Result int64
}

// NewResult Result构造函数
func NewResult(job *Job, result int64) *Result {
	return &Result{
		Job:    job,
		Result: result,
	}
}

// producer 循环生成 int64 类型的随机数，发送到 jobChan
func producer(jobChan chan<- *Job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		jobChan <- NewJob(x)
		time.Sleep(time.Second)
	}
}

// consumer 从 jobChan 中取出随机数，计算各位数的和，将结果发送到 resultChan
func consumer(jobChan <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	for job := range jobChan {
		n := job.Value
		r := calc(n)
		resultChan <- NewResult(job, r)
	}
}

// calc 计算各位数的和
func calc(n int64) int64 {
	s := int64(0)
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func main() {
	var jobChan = make(chan *Job, 100)
	var resultChan = make(chan *Result, 100)

	wg.Add(1)
	go producer(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go consumer(jobChan, resultChan)
	}

	// 主 goroutine 从 resultChan 取出结果并打印到终端
	for r := range resultChan {
		fmt.Printf("value: %19d, result: %3d\n", r.Value, r.Result)
	}

	wg.Wait()
}
