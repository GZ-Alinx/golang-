package main

import (
	"errors"
	"fmt"
)

// Func...
/* Golang 函数 */
// 函数首字母大写，可在包外调用，小写不可调用
func Func() {
	fmt.Println("我是一个函数")
}

// 主函数 位置顺序可以随意放
func main() {
	Func()
	Install(1, 2, 3, 4, 5, 6, 7, 8)
	remove("docker")
	gd("good", "kkkkk")
}

// 函数中参数
func remove(service string) bool {
	if service == "docker" {
		fmt.Println("docker is remove Success~")
	} else if service == "containerd" {
		fmt.Println("failure")
	} else {
		fmt.Println("nil")
	}
	return true
}

// 函数中的不定长参数  返回值
func Install(args ...int) string {
	for _, v := range args {
		fmt.Println(v)
	}
	return "可变参数函数体"
}

// 多参数，多返回值
func gd(args ...string) (string, error) {
	for _, v := range args {
		fmt.Println(v)
	}
	return "args", errors.New("错误返回值")
}
