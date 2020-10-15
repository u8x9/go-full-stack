package main

import (
	"fmt"
	"sync"
	"time"
)

// 在很多编程场景下，需要确保某些操作在高并发的场景下只执行一次
// 例如只加载一次配置文件、只关闭一次 channel

var cfg map[string]interface{}
var loadCfgOnce sync.Once

func loadConfig() {
	cfg = map[string]interface{}{
		"server": "127.0.0.1",
		"port":   9527,
		"ssl":    true,
		"cert":   []string{"path/ssl.crt", "path/ssl.key"},
	}
	fmt.Println("loadConfig invoked")
}

func config(name string) interface{} {
	loadCfgOnce.Do(loadConfig)
	return cfg[name]
}

func main() {
	go func() {
		c := config("host")
		fmt.Println(c)
	}()
	go func() {
		c := config("port")
		fmt.Println(c)
	}()
	go func() {
		c := config("ssl")
		fmt.Println(c)
	}()
	go func() {
		c := config("cert")
		fmt.Println(c)
	}()
	time.Sleep(time.Second)
	fmt.Printf("%#v\n", cfg)
}
