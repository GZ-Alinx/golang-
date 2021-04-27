package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前路径的信息 之心目录
	fmt.Println(".")
	fmt.Println(filepath.Abs("."))
	// 获取文件绝对路劲 二进制文件所在
	fmt.Println(os.Args[0])
	fmt.Println(filepath.Abs(os.Args[0]))
	path, _ := filepath.Abs(os.Args[0])
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))

	// 获取文件路径
	fmt.Println(os.Getwd())
}
