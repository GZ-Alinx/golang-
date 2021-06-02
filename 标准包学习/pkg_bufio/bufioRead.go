package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file,_ := os.Open("bufio.txt")
	reader := bufio.NewReader(file)
	// 直接从bufio缓冲里面调用数据
	data := make([]byte, 3)
	n,_ := reader.Read(data)
	fmt.Println(data[:n])



	//readline 读取文件内容
	ctx, isPerfix,err := reader.ReadLine()
	fmt.Println(string(ctx),isPerfix, err)
}
