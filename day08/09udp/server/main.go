package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed:", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed:", err)
			continue
		}
		fmt.Printf("data: %v addr: %v bytes: %v\n", string(data[:n]), addr, n)
		s := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(s), addr)
		if err != nil {
			fmt.Println("write to udp failed:", err)
			continue
		}
	}
}
