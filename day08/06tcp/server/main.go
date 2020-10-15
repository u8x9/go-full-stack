package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	var buf [128]byte
	for {
		n, err := conn.Read(buf[:])
        if err == io.EOF {
            fmt.Println("See you!")
            break
        }
		if err != nil {
			fmt.Println("read from client failed:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp server listen failed:", err)
		return
	}
	// 2. 等待客户端建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed:", err)
			return
		}
		// 3. 与客户端通信
		go process(conn)
	}
}
