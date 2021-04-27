package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 读取文件使用字节切片进行处理 使用切片中的长度进行读取
	path := "test.txt"

	// file 为一个指针  可以使用go doc os.file 查询使用方法
	file, err := os.Open(path)
	fmt.Println(file, err)
	if err != nil {
		return
	}
	defer file.Close() // 关闭 延迟执行 永远保证关闭句柄

	// 读文件 使用Read() 给它一个切片，把内容放到切片中 然后去提取
	// 所以需要定义一个切片进行存储数据 相当于每次读取的长度
	content := make([]byte, 1024) // 初始化之后为 0
	for {
		// n 读取长度  err 错误
		n, err := file.Read(content) // 读取的结果属于字节
		if err != nil {
			// EOF 标识文件读取结束
			// 非EOF
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(content[:n]))
	}

	// 中文处理 unicode 类型 每一个中文占用3个字节

}
