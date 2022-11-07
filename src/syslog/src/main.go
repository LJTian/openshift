package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == "" {
		port = "30000"
	}
	address := ip + ":" + port
	fmt.Println(address)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	listen, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
