package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("2test.log", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// content := make([]byte, 10) // 切片为读取的字符串大小
	// file.Read(content)
	// fmt.Println(string(content))
	// fmt.Println(file.Write([]byte("123"))) // 写入内容
	// reads := make([]byte, 100)
	// fmt.Println(file.Read(reads))

	//  写入
	file.WriteString("very good")
	text := make([]byte, 1)
	file.Read(text)

	// 移动读取光标
	file.Seek(-10, 6) // 移动到文件最前面
	fmt.Println(file.Read(text))
	fmt.Println(text)
}
