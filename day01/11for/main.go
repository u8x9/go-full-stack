package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------")

	// 变种1
	j := 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("-------")

	// 变种2
	k := 5
	for k < 10 { //for ;k < 10; {
		fmt.Println(k)
		k++
	}

	// 无限循环
	//fmt.Println("-------")
	//for {
	//   fmt.Println("hhhhh")
	//}

	fmt.Println("-------")

	// for range
	// 遍历数组、切片、字符串、map及channel
	s := "Hello沙河"
	for idx, v := range s {
		//fmt.Println(idx, v)
        fmt.Printf("%d => %c\n", idx, v)
	}
}
