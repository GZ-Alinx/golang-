package main

import "fmt"

func main() {
	fmt.Println("start")
	// defer 函数调用
	// defer 延迟执行
	// 在函数退出之前执行
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println("函数退出之前执行", i)
		}()
	}

	fmt.Println("end")
	// 堆栈 谁先声明 谁后执行  这是堆栈

	// 注意
	// 在延迟执行中尽量不要修改返回值

}
