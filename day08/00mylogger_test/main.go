package main

import (
	"time"

	"github.com/u8x9/go-full-stack/day08/mylogger"
)

func main() {
	level := mylogger.DEBUG
	logs := []mylogger.Logger{
		mylogger.NewConsoleLogger(level),
		mylogger.NewFileLogger(level, "hello.log", 1024),
	}
	p := struct {
		name string
		age  int
	}{name: "张三", age: 123}

	c := struct {
		brand string
		price float64
	}{"卡罗拉", 12.34}
	for {
		for _, log := range logs {
			log.Debug("Hello %v", p)
			log.Info("World")
			log.Error("foo bar %v", c)
		}
		time.Sleep(time.Second)
	}
}
