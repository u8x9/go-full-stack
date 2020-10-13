package main

import (
	"fmt"
	"time"

	"github.com/u8x9/go-full-stack/day06/mylogger"
)

type person struct {
	name string
	age  uint8
}

func main() {
	p := person{
		name: "张三",
		age:  123,
	}
	logLevel := mylogger.DEBUG
	cl := mylogger.NewConsoleLogger(logLevel)
	fl := mylogger.NewFileLogger(logLevel, "foo.log", 1024)
    defer fl.Close()
	ls := []mylogger.Logger{cl, fl}
	for {
		for _, l := range ls {
			l.Debug("Hello %v", p)
			l.Trace("World")
			l.Info("你好")
			l.Warning("世界")
			l.Error("foo")
			l.Fatal("bar")
		}
		fmt.Println()
		time.Sleep(time.Second * 3)
	}
}
