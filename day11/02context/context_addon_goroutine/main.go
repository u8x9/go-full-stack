package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要 Context

var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()
    go f2(ctx)
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("隔壁王叔叔")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func f2(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("王叔叔在隔壁")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
