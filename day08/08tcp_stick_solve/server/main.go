package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/u8x9/go-full-stack/day08/08tcp_stick_solve/protocol"
)

const listenAddr string = "127.0.0.1:20000"

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := protocol.Decode(reader)
        if err == io.EOF {
            break
        }
		if err != nil {
			fmt.Println("Decode message failed:", err)
			continue
		}
		fmt.Printf("收到客户端发来的信息：%q\n", msg)
	}
    fmt.Println("See you!")
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}
	fmt.Println("Server listen at", listenAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err)
			return
		}
		fmt.Println("New connection:", conn.RemoteAddr().String())
		go process(conn)
	}
}
