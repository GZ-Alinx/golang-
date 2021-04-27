package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	ctx := []byte("中国最棒")
	hash := fmt.Sprintf("%x \n", md5.Sum(ctx)) // Sprintf 转换为字符串
	fmt.Println(hash)
}
