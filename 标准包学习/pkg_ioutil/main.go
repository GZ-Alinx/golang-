package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// ioutil
	// io

	// 读取文件 所有内容读取进来
	data, _ := ioutil.ReadFile("rdwr.go")
	fmt.Println(string(data))

	// 读取目录 并获取信息  参数： 指定文件路径
	files, _ := ioutil.ReadDir(".")
	for _, info := range files {
		fmt.Println(info.Name())
	}

	// 写文件 参数： 1 文件对象  2 写入内容  3 文件权限
	ioutil.WriteFile("ioutil.txt", []byte("xxx"), os.ModePerm)

	// 创建临时文件 参数 1 路径 2 文件
	path,_ := ioutil.TempFile("./", "a.txt")
	fmt.Println(path)
	//创建临时目录
	paths,_ := ioutil.TempDir("./", "alinx")
	fmt.Println(paths)

}
