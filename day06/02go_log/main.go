package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 标准库里的 log

func main() {
	file, err := os.OpenFile("foo.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	log.SetOutput(file)
	for {
		log.Println("无力吐槽")
		time.Sleep(time.Second * 3)
	}
}
