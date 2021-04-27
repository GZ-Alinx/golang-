package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("bufio.txt")

	//reader := bufio.NewReader(file)
	// 设置缓冲区大小
	reader := bufio.NewReaderSize(file, 5)
	data := make([]byte, 3)
	n, _ := reader.Read(data)
	fmt.Println(data[:n])
	ctx, isPerfix, _ := reader.ReadLine() // 缓冲区中数据处理完成 后若无换行符就会返回 isPerfix为True 需要继续读取
	fmt.Println(string(ctx), isPerfix)

	// 使用\n 进行读取
	fmt.Println(reader.ReadString('\n'))
	fmt.Println(reader.ReadString('\n'))
	fmt.Println(reader.ReadString('\n'))
	fmt.Println(reader.ReadString('\n'))
	fmt.Println(reader.ReadString('\n'))
	fmt.Println(reader.ReadString('\n'))

}
