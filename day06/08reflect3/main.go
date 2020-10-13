package main

import (
	"fmt"
	"reflect"
)

// isNil 通常用于判断指针是否为空
// isValid 通常用于判断返回值是否有效

func main() {
	var a *int
	fmt.Println("var a *int isNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil isValid:", reflect.ValueOf(nil).IsValid())

	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找foo字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(b).FieldByName("foo").IsValid())
    // 尝试从结构体中查找foo方法
    fmt.Println("不存在的结构体方法：", reflect.ValueOf(b).MethodByName("foo").IsValid())

    // map
    m:=map[string]int{}
    // 尝试从map查找一个不存在的键
    fmt.Println("map中不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf("bar")).IsValid())
}
