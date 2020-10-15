package main

import (
	"fmt"
	"net"

	"github.com/u8x9/go-full-stack/day08/08tcp_stick_solve/protocol"
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
		encodingMsg, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("Encode message failed:", err)
			continue
		}
		if _, err = conn.Write(encodingMsg); err != nil {
			fmt.Println("#", i, "write failed:", err)
		}
	}
}
