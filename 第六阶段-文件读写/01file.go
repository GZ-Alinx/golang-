package main

import (
	"fmt"
	"os"
	"time"
)

// go doc os.OpenFile   查看帮助
/*
go doc os.O_CREATE  文件操作模式
package os // import "os"

const (
        // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
        O_RDONLY int = syscall.O_RDONLY // open the file read-only. 只读的方式
        O_WRONLY int = syscall.O_WRONLY // open the file write-only. 只写方式打开
        O_RDWR   int = syscall.O_RDWR   // open the file read-write. 读写方式打开
        // The remaining values may be or'ed in to control behavior.
        O_APPEND int = syscall.O_APPEND // append data to the file when writing. 追加模式
        O_CREATE int = syscall.O_CREAT  // create a new file if none exists. 文件不存在则创建
        O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist. 文件必须不存在
        O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O. 使用同步IO，一般不使用，常使用异步IO
        O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened. 打开截断文件
		截断文件： 清空

*/

func main() {
	// 打开文件 并给0600权限
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return
	}
	// 延迟关闭
	defer file.Close()
	// 写入时间到文件中
	fmt.Fprintf(file, "%s\n", time.Now().Format("2006-01-02 15:04:05")) // 写入时间到文件中
}

// 文件方法实现
// func Open(name string) (*File, error) {
// 	os.OpenFile(name, os.O_RDONLY, 0777)
// }
// func Create(name string) (*File, error) {
// 	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
// }
