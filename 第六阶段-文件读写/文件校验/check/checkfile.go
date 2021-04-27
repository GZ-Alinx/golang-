package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const BufferSize = 1024

// 文件如果太大 可以切片文件内容逐个进行校验
func main() {
	var (
		src string
		dst string
	)

	flag.StringVar(&src, "src", "", "src file path")
	flag.StringVar(&dst, "dst", "", "dst file path")

	flag.Usage = func() {
		fmt.Println("usage: cp -src path --dst path")
	}

	flag.Parse()

	// 检查
	/*
		src
			没有输入
			文件路径不存在
			目录

		dst
			没输入
			目录不存在 => 自动创建父目录（创建失败/创建成功，或者报错

			文件存在
				是目录
	*/

	// 打开文件
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延迟关闭
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()
	ctx := make([]byte, BufferSize)

	// 读取源文件，写入目的文件
	for {
		n, err := srcFile.Read(ctx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		// 写入
		dstFile.Write(ctx[:n])
	}
}
