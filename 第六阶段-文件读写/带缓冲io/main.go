package main

import (
	"bufio"
	"fmt"
	"os"
)

// 带缓冲IO
//	程序中定义切片，内部定义切片长度，读取内容到切片，
//	函数直接从切片中读取数据，当切片中的数据处理完成后，再次冲硬盘读取
//	，写数据到切片， 当写满后刷新到磁盘（如果没有写满如何刷新到磁盘）

// 带缓存区IO 读取的数据为字符串   strconrv转换
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 实时输入输出
	for scanner.Scan() {
		fmt.Println("输入内容:", scanner.Text(), scanner.Bytes())
	}
}
