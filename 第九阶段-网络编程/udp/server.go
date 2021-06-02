package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	//  UDP监听包
	addr := ":9999"
	server,err :=  net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// 延迟关闭
	defer server.Close()
	for {
		// 接收数据包（发送者信息addr， 发送内容）
		buffer := make([]byte, 1024) // 定义一个容器来对数据进行操作
		n,addrs,err := server.ReadFrom (buffer)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("recv:", addrs,string(buffer[:n]))
		}
	}

}
