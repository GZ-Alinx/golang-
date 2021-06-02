package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// 和tcp一样，只是协议需要写成udp
	conn,err := net.Dial("udp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(conn)
	}
	defer conn.Close()
	// 给服务器端写了一个时间
	for {
		go conn.Write([]byte(time.Now().Format("2006-01-02 15:03:04")))
		time.Sleep(111)
	}

}