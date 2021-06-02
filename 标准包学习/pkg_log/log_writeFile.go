package main

import (
	"log"
	"os"
)

func main() {
	// 将日志记录到文件中
	// 后面参数： 文件名称  打开模式 文件权限
	file,err := os.OpenFile("web1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file) //设置日志写入文件
	log.Print("我是日志")
	log.Println("换行日志。。。。")
	log.Println("换行日志。。。。")
	log.Println("换行日志。。。。")
	log.Println("换行日志。。。。")
}
