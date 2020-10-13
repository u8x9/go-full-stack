package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f() {
	for i := 0; i < 5; i++ {
		i1 := rand.Int()
		i2 := rand.Intn(10) // 0~9
		fmt.Println(i1, i2)
	}
}

func f1(i int) {
	defer wg.Done() // 3. 减少计数
	time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
	fmt.Println("goroutine No.", i)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 1. 增加计数
		go f1(i)
	}
	wg.Wait() // 2. 等待计数结束
	fmt.Println("main")
}
