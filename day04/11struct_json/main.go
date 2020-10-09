package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "张三",
		Age:  123,
	}
	var buffer []byte
	var err error
	if buffer, err = json.Marshal(p); err != nil {
		fmt.Println("json序列化失败！")
		return
	}
	fmt.Println(string(buffer))

	var p1 Person
	if err = json.Unmarshal(buffer, &p1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", p1)
}
