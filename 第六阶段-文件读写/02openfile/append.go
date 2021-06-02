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

	//pkg_fmt.Sprint("打印")
	//pkg_fmt.Sprintf("组合打印")
	//pkg_fmt.Sprintln("")
	//
	//files, err := os.Create("alinx.log")
	//pkg_fmt.Fprint(files, 1,2,3)
	//pkg_fmt.Fprint(files, "x","y","z")
	//pkg_fmt.Fprintln(files,4,5,6)

}
