package main

import (
	"fmt"
	"io/fs"
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



	// 正则 写法
	filepath.Glob("./*.go") // 所有以.go结尾的文件
	filepath.Glob("./*.txt")
	// 多种格式 （f1|f2|f3） ？匹配一个  * 匹配任意多个
	fmt.Println(filepath.Glob("./*.go|./*.exe"))

	// 路径匹配
	fmt.Println(filepath.Match("a/*.go", "a/1.go"))



	// WalkDir 递归目录 拿到文件信息
	// 参数  目录  函数（返回路径， 返回文件信息， 返回错误信息）
	filepath.Walk("../", func(path string, info fs.FileInfo, err error) error {

		fmt.Println(path)
		return nil
	})
}
