package main

import (
	"fmt"
	"sync"
)

// 互斥锁

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 1; i <= 100; i++ {
		lock.Lock() // 加锁
		x += i
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Printf("x = %d\n", x) // 505000
}
