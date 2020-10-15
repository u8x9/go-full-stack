package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const listenAddr string = "127.0.0.1:20000"

func process(conn net.Conn) {
	defer conn.Close()
    reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			fmt.Println("See you!")
			break
		}
		if err != nil {
			fmt.Println("Read from client failed:", err)
			continue
		}
		fmt.Printf("收到客户端发来的信息：%q\n",string(buf[:n]))
	}
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
