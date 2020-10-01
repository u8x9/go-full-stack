package main

import "fmt"

// 流程控制之跳出for循环

func main() {
	// 当i==5 时，跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	fmt.Println("====")
	// 当 i == 5 时，跳出本次循环，继续下次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
