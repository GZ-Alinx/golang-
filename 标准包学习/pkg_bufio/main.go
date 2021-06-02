
package main

import (
"bufio"
"fmt"
"os"
)

// 实时输出 标准输入的数据

// 带缓冲的io
//读取数据 再strconv转换数据类型


/*
	程序中定义类似切片，切片长度， 读取内容到切片 函数直接从切片中读取数据，当切片中的数据处理完成后，
再从硬盘读取
	写数据到切片， 当写满后刷新到磁盘 如果没有写满， 如何刷新到磁盘？ 如果没有刷新数据是否丢了？

	bufio
	Scanner
	Reader
	Writer
*/

// scanner
func Scanner() {
	// 实时扫描输入的数据 进行处理
	scanner := bufio.NewScanner(os.Stdin)
	//var sum int
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// 返回的数据类型不同   以换行为标识符
		fmt.Println("输入内容:",scanner.Text(), scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}



func main() {
	Scanner()
}

