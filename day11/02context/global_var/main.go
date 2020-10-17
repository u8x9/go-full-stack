package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要 Context
// 如何通知子 goroutine 退出？

var wg sync.WaitGroup
var notify bool

func f1() {
	defer wg.Done()
	for {
		fmt.Println("隔壁王叔叔")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
