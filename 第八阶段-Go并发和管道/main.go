package main

import (
	"fmt"
)


func test() {
	fmt.Println("I am work in a single goroutine")
}

func main() {
	fmt.Println("并发学习")
	go test()
	//time.Sleep(time.Second) 如果没有time 那test 基本不会执行 因为main主进程已经结束了 这样就以为着整个进程运行结束了

}
