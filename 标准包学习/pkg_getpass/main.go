package main

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/princebot/getpass"
)

// 密码加密

func main() {
	data, _ := getpass.Get("dfad:")
	fmt.Printf("%v,%q", data, string(data))

	datas, _ := gopass.GetPasswd()
	fmt.Printf("%v,%w", datas, string(datas))
	// 输出的属于字节 可以进行转换为字符串
}
