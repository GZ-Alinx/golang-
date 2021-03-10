package main

import "fmt"

/*
	main 函数 程序入口
	开发者写代码的执行入口
*/

func main() {
	fmt.Println("程序入口")
}

// 初始化函数
// 导入包时执行
func init() {
	fmt.Println("pkg.init")
}

/*
导入包流程
1. main 函数加载 从上到下 加载程序
2. 发现import包 寻找import中的包 并加载所有go文件
	加载go文件 检查是否用init函数，有就执行
	同时加载变量，常量等

	以此循环包加载


	注意！！！！！！！！！！
	go中标识符只能定义一次 但是init例外
	go文件中可以包含多个  init函数

	最后执行main
*/
