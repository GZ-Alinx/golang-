package main

import "fmt"

// 函数在调用是会在堆栈中开辟内存空间，进行存放数据，函数调用结束后会释放
// 匿名函数
// lambada 匿名函数

func main() {
	// names := []string{"alinx", "nginx", "linux"}
	// 匿名函数定义
	lambada := func() {
		fmt.Println("我是一个匿名函数")
	}
	fmt.Printf("%T ,%q\n", lambada, lambada)
	// 或者如下调用
	lambada()

	// 匿名函数直接调用
	func() {
		fmt.Println("匿名函数直接调用 在函数末尾中\"()\" 引用")
	}()
}
