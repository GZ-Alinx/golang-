package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := os.Args[1]
	if len(os.Args) != 2 {
		log.Fatal("参数不正确")
	}
	files, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer files.Close()
	info, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("是否为目录:", info.IsDir())
	fmt.Println("文件名称:", info.Name())
	fmt.Println("文件权限:", info.Mode())
	fmt.Println("文件修改时间:", info.ModTime())
	fmt.Println("文件大小:", info.Size())
	fmt.Println("系统信息:", info.Sys())

	fmt.Println("时间类型修改为字符串类型：", info.ModTime().Format("2006-01-02 15:03:04"))

	//hasher := md5.New()
	//hasher.Write(os.Args[1])

}

func manyArgs() (string, string) {
	a := "123"
	b := "321"
	return a, b
}
