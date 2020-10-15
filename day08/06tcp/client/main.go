package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed:", err)
		return
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("# Please enter message:")
		msg, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("get user input failed:", err)
		}
		msg = strings.TrimSpace(msg)
		if msg == ".exit" {
			fmt.Println("Bye!")
			break
		}
		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("# write to server failed:", err)
			return
		}
		fmt.Println("# success send", n, "bytes data to server")
		fmt.Println()
	}
}
