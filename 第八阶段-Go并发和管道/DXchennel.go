package main

import "fmt"

// 默认清空下 通道是双向的 也就是可以从里面写数据 也可以从哪里拿数据

// 但是channel是可以让它只能读  或者 只能写的 操作 来对应我们的需求

//单向申明
var ch1 chan int       // 正常申明 普通channel
var ch2 chan<- float64 // 单向 只可写 fload64类型的数据
var ch3 <-chan int     // 单项只能用于读的管道

//可以将channel隐式的转换为单项队列，只收或只发， 不能将单项的channel转换为普通channel

// 管道写入
func product(in chan<- int) {
	for i := 0; i <= 10; i++ {
		in <- i
	}
}

// 管道写出
func comsumer(out <-chan int) {
	for name := range out {
		fmt.Println(name)
	}
}

// 读取管道中的数据  单项读 或者单向写
func main() {
	ch := make(chan int)
	go product(ch)
	comsumer(ch)
}
