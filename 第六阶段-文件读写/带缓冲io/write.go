package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("write.txt")
	//write := bufio.NewWriter(file)
	write := bufio.NewWriterSize(file, 10)
	fmt.Println(write.WriteString("99999999999"))

	//	刷新数据写入文件
	write.Flush()
}
