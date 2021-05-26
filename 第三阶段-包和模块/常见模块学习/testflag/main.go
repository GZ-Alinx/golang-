package main

import (
	"flag"
	"fmt"
)

func main() {
	// 先创建变量
	// ssh -P port root@host
	var port int
	var help bool
	var password string
	// 指定函数和命令行参数的关系
	/*
		1 &port   参数绑定
		2 "P"     命令行中参数名称
		3 22      默认值
		4 "port"  帮助信息
	*/
	flag.IntVar(&port, "P", 22, "port")
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&password, "p", "", "password")

	flag.Usage = func() {
		fmt.Println("usage: ssh -P 22 root@host")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	// 帮助信息
	if help {
		fmt.Println("帮助信息")
		flag.PrintDefaults() // 帮助信息调取
		return
	}
	fmt.Println(port)

	// 使用
	// windows项目 run .\main.windows项目 -P

	// 打印没有指定的参数名称
	fmt.Println(flag.Args())
}
