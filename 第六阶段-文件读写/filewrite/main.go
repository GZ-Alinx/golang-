package main

import (
	"fmt"
	"os"
	"time"
)

// 打开文件标识 flag
/*
go doc os.O_CREATE
	O_RDONLY 	以只读的方式打开 		open the file read-only.
	O_WRONLY 	以只写的方式打开 		open the file write-only.
	O_RDWR   	以读写的方式打开 		open the file read-write.
	O_APPEND 	追加            		append data to the file when writing.
	O_CREATE 	文件不存在则创建 		create a new file if none exists.
	O_EXCL   	文件必须不存在  		used with O_CREATE, file must not exist.
	O_SYNC   	使用同步IO      		open for synchronous I/O.
	O_TRUNC  	截断文件 相当于清空      truncate regular writable file when opened.

*/

// 属性组合 ，成为自己的方法
func Open(name string) (*File, error) {
	return os.OpenFile(name, os.O_RDONLY, 0777)
}

func Create(name string) (*File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0777)
}

func main() {
	// 打开文件标识 使用或运算累加
	/*
		第一个参数： 文件路径
		第二个参数： 控制文件打开方式
		第三个参数： 控制文件模式
	*/

	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	// 另外一种写法
	// os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	fmt.Fprintf(file, "%s\n", time.Now().Format("2006-01-02 15:04:05"))
}
