package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要 Context

var wg sync.WaitGroup

func f1(ch <-chan struct{}) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ch:
			break LOOP
		default:
			fmt.Println("隔壁王叔叔")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var notify = make(chan struct{})
	wg.Add(1)
	go f1(notify)
	time.Sleep(time.Second * 5)
	notify <- struct{}{}
	close(notify)
	wg.Wait()
}
