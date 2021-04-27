package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(filepath.Abs("."))

	// 获取文件路径 文件名称
	path := "alinx.txt"
	fmt.Println(filepath.Dir(path), filepath.Base(path), filepath.Ext(path))

	filepath.Clean("./a/")
	fmt.Println(filepath.Join("./", "a")) // 路径拼接
	strings.Split("1 2 3 ", " ")          // 字符分隔 用空格分割前一个参数的字符
	filepath.Split(".")                   // 返回目录和文件名称
}
