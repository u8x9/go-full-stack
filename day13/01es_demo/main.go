package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("create elastic client failed:", err)
		return
	}
	p := Person{
		Name:    "张三丰",
		Age:     999,
		Married: false,
	}
	putedResult, err := client.Index().Index("user").BodyJson(p).Do(context.Background())
	if err != nil {
		fmt.Println("put failed:", err)
	}
	fmt.Printf("indexed user: %s to index %s,type: %s\n", putedResult.Id, putedResult.Index, putedResult.Type)
}
