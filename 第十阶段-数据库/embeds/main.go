package main

import (
	"embed"
	_ "embed"
	"fmt"
)
//将所有的文件编译到二进制文件中

// 版本管理  embed




// 定义使用  只能再当前目录中定义version文件
//go:embed version.txt
//go:embed name.txt



// 多个文件打包到二进制中 整个目录 或者某种格式文件 进行包含
//go:embed a
var fs embed.FS





func main() {
	// 多个文件包含到二进制文件中
	file,err := fs.ReadFile("a/name")
	if err != nil  {
		fmt.Println(err)
	}
	fmt.Println(string(file))
	file,err = fs.ReadFile("a/set")
	fmt.Println(string(file))




}
