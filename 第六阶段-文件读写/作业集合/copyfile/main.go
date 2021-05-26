package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

/*
1. copyfile 测试bufferSize大小对程序执行速度影响
time copyfile

time.Now()
Sub()
*/

func main() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
	// 计算整体复制时间 毫秒
	startTime := time.Now().UnixNano()

	// 计算args参数个数  少于3则不符合
	if len(os.Args) != 3 {
		log.Fatal("参数不正确")
	}

	srcFile, dstFile := os.Args[1], os.Args[2]

	// 打开需要拷贝的文件
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	// 延迟关闭
	defer src.Close()

	// 拷贝目标文件创建
	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	buffsize := 1

	// 创建存放数据的切片
	data := make([]byte, buffsize)
	// 循环打开文件进行读取和写入
	for {
		// 读取源文件数据
		n, err := src.Read(data)
		if err != nil {
			if err == io.EOF {
				endTime := time.Now().UnixNano() // 毫秒时间
				fmt.Println("使用时间:", (endTime-startTime)/1e6)
				//fmt.Println("使用时间:",(endTime - startTime) / 1e9)
				log.Fatal(err)
			}
			//fmt.Println("文件读取完毕")
			break
		}

		// 写入数据到 目标文件
		n, err = dst.Write(data[:n])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("文件写入完毕")
	}
}