package main

import "fmt"

func main() {
	fmt.Println("start")
	// defer 函数调用
	// defer 延迟执行
	// 在函数退出之前执行
	defer func() {
		fmt.Println("函数退出之前执行 1")
	}()
	defer func() {
		fmt.Println("函数退出之前执行 2")
	}()
	fmt.Println("end")
	// 堆栈 谁先声明 谁后执行  这是堆栈

}
