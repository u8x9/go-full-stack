package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("dial to server failed:", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	var msg string
	for {
		fmt.Println("Please enter message(`.exit` to quit):")
		msg, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("read user input failed:", err)
			continue
		}
		msg = strings.TrimSpace(msg)
		if msg == ".exit" {
			break
		}
		sendData := []byte(msg)
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Println("send data failed:", err)
			continue
		}
		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read data failed:", err)
			continue
		}
		fmt.Printf("recv: %v, addr: %v, bytes: %v\n", string(data[:n]), remoteAddr, n)
		time.Sleep(time.Second)
	}
	fmt.Println("Bye!")
}
