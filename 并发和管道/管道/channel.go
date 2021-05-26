package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// 历程
	go func() {
		// 关闭管道
		time.Sleep(3 * time.Second)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

		v,ok := <-channel
		fmt.Println(v,ok)
	}()


	// 不能针对关闭的管道写数据
	// 关闭的管道还能写 还能读 吗

	channel <- 1
	close(channel)
	time.Sleep(3 * time.Second)

}
