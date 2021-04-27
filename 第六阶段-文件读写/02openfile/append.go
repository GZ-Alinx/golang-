package main

import (
	"log"
	"os"
)

// 标准文件读写
func main() {
	path := "web.log"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Fatal("第一条日志")

	//fmt.Sprint("打印")
	//fmt.Sprintf("组合打印")
	//fmt.Sprintln("")
	//
	//files, err := os.Create("alinx.log")
	//fmt.Fprint(files, 1,2,3)
	//fmt.Fprint(files, "x","y","z")
	//fmt.Fprintln(files,4,5,6)

}
