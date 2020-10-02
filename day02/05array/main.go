package main

import "fmt"

// 数组
// 数组的长度是类型的一部分

func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("%T, %T\n", a1, a2)

	// 数组的初始化
	// 如果没有初始化，所有元素都是零值
	fmt.Println(a1, a2)
	// 初始化方式1
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)
	// 初始化方式2
	// 根据初始化值自动推断数组的长度
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("%T, %v\n", a3, a3)
	// 初始化方式3
	a4 := [5]int{1, 2}
	fmt.Printf("%T, %v\n", a4, a4)
	// 初始化方式4(指定索引来初始化)
	a5 := [5]int{0: 1, 4: 2}
	fmt.Printf("%T, %v\n", a5, a5)

	// 数组的遍历
	citys := [...]string{"上海", "广州", "深圳", "澳门"}
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Printf("citys[%d] = %q\n", i, citys[i])
	}
	// 2. for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	a6 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(a6)

	// 数组是值类型
	a7 := [3]int{1, 2, 3}
	a8 := a7
	a8[1] = 999
	fmt.Println(a7, a8)
}
