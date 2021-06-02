package main

import (
	"fmt"
	"os"
)

func main() {

	/* scan
	fmt.Scan() 标准输入上扫描
	fmt.Sscan() 字符串中扫描
	fmt.Fscan() ioReader中扫描
	*/



	/*
	控制台，文件擦走
	标准输入，标准暑促和 标准错误输出
	fmt.Print()

	*/
	// fmt Fprint数据流写入 将右边的数据写入左边的对象中
	file, _ := os.Create("fprintf.txt")
	//数据流写入
	fmt.Fprint(file, 1,2,3)
	fmt.Fprint(file, "x")
	fmt.Fprintln(file, "x","y","z")
	fmt.Fprintf(file, "%s", "123")
}
