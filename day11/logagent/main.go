package main

import (
	"fmt"
	"time"

	"github.com/u8x9/go-full-stack/day11/logagent/kafka"
	"github.com/u8x9/go-full-stack/day11/logagent/taillog"
)

func run() {
	for {
		select {
		// 1. 读取日志
		case line, ok := <-taillog.Read():
			if ok {
				// 2. 发送到kafka
				pid, offset, err := kafka.SendMessage(line.Text)
				if err != nil {
					fmt.Println("SendMessage failed:", err)
					break
				}
				fmt.Println("send to pid:", pid, ",offset:", offset)
			} else {
				fmt.Println("read log failed")
			}
		default:
            time.Sleep(time.Second)

		}
	}
}
func main() {
	// 1. 初始化kafka连接
	if err := kafka.InitDefaultLocalSingle(); err != nil {
		panic(err)
	}
	defer kafka.Close()
	// 2. 打开日志文件准备收集日志
	if err := taillog.Init("./my.log"); err != nil {
		panic(err)
	}
	// 3. 运行
	run()
}
