package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x      = 0
	wg     sync.WaitGroup
	rwlock sync.RWMutex
	lock   sync.Mutex
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	fmt.Println(x)
	//lock.Unlock()
	rwlock.RUnlock()
	time.Sleep(time.Millisecond)
}
func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	x += 1
	//lock.Unlock()
	rwlock.Unlock()
	time.Sleep(time.Millisecond * 5)
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println("x =", x)
	fmt.Println(time.Now().Sub(start))
}
