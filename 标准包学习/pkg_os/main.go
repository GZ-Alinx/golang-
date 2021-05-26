package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Getwd())        // 查看当前工作目录
	fmt.Println(os.Chdir("alinx")) // 工作目录切换
	fmt.Println(os.Getwd())
	fmt.Println(os.Chmod("windows项目.mod", 0777)) // 更改文件权限
	fmt.Println(os.Chown("windows项目.mod", 1, 1)) // 文件属组和属主变更 已id号为准 只支持Linux系统

	fmt.Println(os.Create("xxx.txt")) // 创建文件
	// 位置参数
	fmt.Println(os.Args) // args 参数

	// 文件访问时间和修改时间 修改
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	fmt.Println(os.Chtimes("windows项目.mod", atime, mtime)) // 修改文件的访问时间和修改时间
	// 删除所有的ENV： os.Clearenv()
	fmt.Println(os.DirFS("/")) // 返回目录的文件系统
	// 退出码
	// os.Exit(0) // 0表示成功， 1表示失败

	fmt.Println(os.Getenv("GOPATH"))                   // 获取ENV环境信息 没有则返回空
	fmt.Println(os.Hostname())                         // 返回主机名
	fmt.Println(os.Lchown("windows项目.mod", 1, 1))             // 更改文件的uid 和 gid 不支持windows
	fmt.Println(os.Mkdir("test", 0777))                // 创建目录并赋予权限
	fmt.Println(os.MkdirAll("test/123/123/123", 0777)) // 创建目录并赋予权限 类似与-p 递归创建
	fmt.Println(os.Remove("test/123/123/123"))         // 删除目录
	fmt.Println(os.RemoveAll("test"))                  // 递归删除目录
	fmt.Println(os.Chdir("../"))                       // 工作目录切换

	fmt.Println(os.Rename("file.txt", "lixiong")) // 文件改名
	fmt.Println(os.Rename("lixiong", "file.txt")) // 文件改名

	fmt.Println(os.Setenv("alinx", "feichangbang"))     // 设置env环境变量 都为string类型
	fmt.Println(os.Symlink("file.txt", "testfile.txt")) // 软连接创建

	fmt.Println(os.UserHomeDir()) // 返回当前用户主目录
	// os.ReadDir("./")  查看当前目录 类似ls
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())         // 输出文件名称列表
		fmt.Println(file.IsDir())        // 是否是目录
		fmt.Println(file.Type())         // 类型信息
		fmt.Println(file.Type().IsDir()) // 是否是目录一致
	}

}
