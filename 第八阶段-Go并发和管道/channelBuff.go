package main

import (
	"fmt"
	"time"
)

// 如果我们创建channel时制定了channel的长度那么channel将会拥有缓冲区
// goroutine在缓冲区为满时往channel发送数据将不会被阻塞

func consue(ch chan int) {
	// 进程休息10秒 再从哪个ch中读取数据
	time.Sleep(time.Second * 10)
	<- ch
}

func main() {
	// 创建一个长度为2的channel
	ch := make(chan int, 2)
	go consue(ch)

	//发送和接收数据的数量要和定义式一直 否则会阻塞

	// 往channel中发送数据
	ch <- 0
	ch <- 1
	//发送数据不会被阻塞
	fmt.Println("数据发送完毕")
	ch <- 2 // 发送的数据 大于定义的2个数量 所以阻塞了

	//// 读取数据
	//a := <- ch
	//b := <- ch
	//c := <- ch
	//pkg_fmt.Println(a,b,c)
	time.Sleep(time.Second)

}

