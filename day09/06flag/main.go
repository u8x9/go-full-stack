package main

import (
	"flag"
	"fmt"
	"time"
)

func f1() {
	// 创建标志位参数
	name := flag.String("name", "张三", "指定你的名字")
	age := flag.Int("age", 20, "指定你的年龄")
	isMarriage := flag.Bool("marriage", false, "指定婚否")
	countDown := flag.Duration("cd", time.Second, "倒计时")

	// 解析标志位
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*isMarriage)
	fmt.Println(*countDown)
}
func main() {
	var name string
	var age int
	flag.StringVar(&name, "name", "张三", "指定你的名字")
	flag.IntVar(&age, "age", 123, "指定你的年龄")

	flag.Parse()

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("flag外所有参数：", flag.Args())
	fmt.Println("flag外参数个数：", flag.NArg())
	fmt.Println("flag个数:", flag.NFlag())
	/*
	   $ run
	   go run main.go -name lisi  foo bar xixi haha

	   $ output:
	   lisi
	   123
	   flag外所有参数： [foo bar xixi haha]
	   flag外参数个数： 4
	   flag个数: 1
	*/
}
