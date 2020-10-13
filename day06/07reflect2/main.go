package main

import (
	"fmt"
	"reflect"
)

// 通过反射设置变量的值

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，会引发 panic: reflect.Value.SetInt using unaddressable value
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem() 方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	i := int64(123)
	//reflectSetValue(i)
	reflectSetValue2(&i)
	fmt.Println(i)
}
