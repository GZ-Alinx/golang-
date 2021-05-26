package main

import "fmt"

// channel 作为goroutine之间同步和通信的手段 相当于列队
// channel 属于引用类型  每一个channel只能传递固定类型的数据

// 关键字 chan 定义channel
func main() {
	// 声明一个新的channel 并指定channel内传输的数据类型T
	var channelName chan string
	fmt.Printf("%T,  %#v\n", channelName,channelName)

	ch := make(chan string, 1)
	// channel的发送和接收
	//数据总是会遵循先进先出的原则进行 同时也保证同一时刻内仅有一个goroutine访问channel来发送和获取数据

	//从channel发送数据需要使用 "<-" 符号，如下所示
	a := "123"
	ch <- a // 讲a发送到channel中 在channel被填满之后在向通道中发出将会阻塞当前goroutine
	//接收数据
	a,ok := <- ch
	if ok {
		fmt.Println(a)
	}else{
		fmt.Println("数据为空")
	}

}

