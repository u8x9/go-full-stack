package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

var wg = new(sync.WaitGroup)
var lock = new(sync.Mutex)

func add() {
	defer wg.Done()
	//lock.Lock()
	//x++
	//lock.Unlock()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
