package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("INVALID TRACE CODE")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code: %s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
}
func main() {
	// 50ms超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置 trace code 传递给后续启动的 goroutine
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "134813849")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
