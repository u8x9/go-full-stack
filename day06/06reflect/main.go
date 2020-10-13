package main

import (
	"fmt"
	"reflect"
)

type cat struct{}
type u8 uint8
type i32 = int32

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type name: %v, kind: %v\n", t.Name(), t.Kind())
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int64:
		// v.Int() 从反射中获取整型的原始值，然后通过int64()强转
		fmt.Printf("type is int64, value is: %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	case reflect.String:
		fmt.Printf("type is string, value is %q\n", v.String())
	default:
		fmt.Printf("i don't known how to cast it: %#v\n", v.Interface())
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	reflectType(&a)
    reflectValue(a)

	var b int64 = 100
	reflectType(b)
	reflectType(&b)
    reflectValue(b)

	c := cat{}
	reflectType(c)
	reflectType(&c)
    reflectValue(c)

	i := u8(123)
	reflectType(i)
	reflectType(&i)
    reflectValue(i)

	var j i32 = 456
	reflectType(j)
	reflectType(&j)
    reflectValue(j)
}
