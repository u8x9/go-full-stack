package main

import "fmt"

// channel 练习

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0-99发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- (i * i)
		}
		close(ch2)
	}()

	// 在主goroutine中，从ch2中接收数据
	for i := range ch2 {
		fmt.Println(i)
	}
}
