package main

import (
	"fmt"
	"net"
)

const serverIP = "127.0.0.1:20000"

func main() {
	conn, err := net.Dial("tcp", serverIP)
	if err != nil {
		fmt.Printf("Dial to server %q failed: %v\n", serverIP, err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("你好呀Hello%02d", i)
		if _, err = conn.Write([]byte(msg)); err != nil {
			fmt.Println("#", i, "write failed:", err)
		}
	}
}
