package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)


// 浏览器测试

func handle(conn net.Conn) {
	defer conn.Close()
	// 本地地址和 远程链接地址
	fmt.Println("client connect addr:", conn.LocalAddr())
	//pkg_fmt.Println("addr:", conn.RemoteAddr())


	//超时时间设置 按照顺序配置

	// Fprintf 给流里面写数据
	fmt.Fprintf(conn,time.Now().Format("2006-01-02 15:04:05"))

	for {
		sizeBUffer := make([]byte, 5)
		n, err := conn.Read(sizeBUffer)
		if err == nil {
			if n == 5 {
				size,err := strconv.Atoi(string(sizeBUffer))
				if err == nil {
					contentBuffer := make([]byte, size)
					n,err := conn.Read(contentBuffer)
					if err == nil && n == size {
						fmt.Println(contentBuffer)
					}
				}
			}
		}
	}
}

func main() {
	addr := "127.0.0.1:8888"
	listener,err := net.Listen("tcp",addr)
	if err != nil {
		log.Fatalln(err)
	}


	// 管道进行处理
	//interrupt := make(chan int, 1)
	//
	//END:
	//for {
	//	select {
	//	case <-interrupt:
	//		break END
	//	default:
	//
	//	}
	//}
	fmt.Println("listene:",listener.Addr())
	defer listener.Close()

	// 接收客户端请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accept: %s", err)
			break
		}
		go func() {
			//关闭客户端链接
			// 逻辑处理
			// 历程
			go handle(conn)
		}()
	}




}
