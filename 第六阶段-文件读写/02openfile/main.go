package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 文件打开操作属性

// go doc os.OpenFile   查看帮助
/*
go doc os.O_CREATE  文件操作模式
package os // import "os"

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only. 只读的方式
	O_WRONLY int = syscall.O_WRONLY // open the file write-only. 只写方式打开
	O_RDWR   int = syscall.O_RDWR   // open the file read-write. 读写方式打开
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing. 追加模式
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists. 文件不存在则创建
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist. 文件必须不存在
	O_SYNC   int = syscall.O_SYN  C   // open for synchronous I/O. 使用同步IO，一般不使用，常使用异步IO
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened. 打开截断文件
	截断文件： 清空

*/

//创建文件并写入数据
func main() {
	path := "rdwr.txt"
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC // 文件不存在创建,TRUNC 清空文件
	file, err := os.OpenFile(path, flags, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("success")

	//写入文件内容 字符串数据
	file.WriteString("ABC")

	// 文件偏移量 从哪里开始写数据
	fmt.Println(file.Seek(3, 3)) // 修改文件光标的位置

	// 设置切片进行数据存储 进行读取
	data := make([]byte, 4)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
			}
			log.Fatal(err)
		}
		fmt.Println(data, n, string(data[:n]))
	}
}
