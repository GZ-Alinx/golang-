package main

import (
	"flag"
	"fmt"
)

func main() {

	//用于开发这样的指令 ssh -H 127.0.0.1 --port=8080 -v touch xxxx

	// 定义变量
	// 定义命令行格式
	// 元素 => 解析到变量， 默认值，帮助信息

	var (
		host    string
		port    int
		verbose bool
		help    bool
	)
	// 定义具体参数
	flag.StringVar(&host, "H", "127.0.0.1", "连接端口")
	flag.IntVar(&port, "port", 8080, "连接端口")
	flag.BoolVar(&verbose, "v", false, "显示详情")
	flag.BoolVar(&help, "h", false, "帮助信息")

	// 使用帮助
	flag.Usage = func() {
		fmt.Println("usage: ssh -H 127.0.0.1 --port=8080 -v touch xxxx")
		flag.PrintDefaults()
	}

	// 解析
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	fmt.Println(host, port, verbose, help, flag.Args())

}
