package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888" // 链接服务端的地址
	conn,err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1023)
	n,err := conn.Read(buffer)
	if err == nil {
		fmt.Println(string(buffer[:n]))
	}
}
