package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func putAndGet() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("new clientv3 failed:", err)
		return
	}
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "姓名", "张三")
	cancel()
	if err != nil {
		fmt.Println("put failed:", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "姓名")
	cancel()
	if err != nil {
		fmt.Println("get failed:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s = %s\n", ev.Key, ev.Value)
	}
}

func watch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("new failed:", err)
		return
	}
	defer cli.Close()

	// watch key:姓名 change
	rch := cli.Watch(context.Background(), "姓名") // <-chan watchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s, Key: %s, Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func main() {
	//putAndGet()
	watch()
}
