package main

import "fmt"

// 直接抛出错误
func main() {
	defer func() {
		// recover 执行是读取 panic 的错误信息 如果没有panic 则不会执行
		if err := recover(); err != nil {
			fmt.Println("出错了，进行恢复", err)
		}
	}
	// 存在错误，可使用recover进行捕获 如果没有recover则会向往抛出到解释器 即出错
	panic("我1是错误") 
}
