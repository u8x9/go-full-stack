package main

import "fmt"

// 关闭 channel

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1) // 关闭之后依然可以从channel取值
	//for i:=range ch1 {
	//fmt.Println(i)
	//}
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)

	//ch1<-30 // 关闭之后不可以再发送数据：send on closed channel

}
