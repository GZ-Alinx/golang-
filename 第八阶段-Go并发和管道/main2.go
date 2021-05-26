package main

import (
	"fmt"
	"time"
)

func setVto1(v *int) {
	*v = 1
}

func setVTo2(v *int) {
	*v = 2
}


func main() {
	v := new(int)
	fmt.Println(*v)
	go setVto1(v)
	go setVTo2(v)

	// 匿名函数 使用go 关键字启动并发
	go func() {
		fmt.Println("我是一个匿名函数，被go并发goroutine 启动运行")
	}()
	fmt.Println(*v)
	time.Sleep(time.Second)
}
