package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	s := student{
		Name:  "张三",
		Score: 92,
	}

	t := reflect.TypeOf(s)
	fmt.Println("type name:", t.Name(), ", type kind:", t.Kind())

	// 遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("[FIELD] name: %s, index: %d, type: %v, json tag: %v\n", field.Name, field.Index, field.Type.Name(), field.Tag.Get("json"))
	}

    // 通过字段名获取指定结构体字段信息
    if scoreField, ok:=t.FieldByName("Score"); ok {
        fmt.Printf("name: %s, index: %d, type: %v, json tag: %v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
    }
}
