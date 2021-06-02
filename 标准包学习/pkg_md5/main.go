package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	// md5 加密
	// 一次性计算
	//ctx := []byte("中国最棒")
	//hash := pkg_fmt.Sprintf("%x \n", md5.Sum(ctx)) // Sprintf 转换为字符串
	//pkg_fmt.Println(hash)

	// 累加的方法进行计算
	hasher := md5.New()
	hasher.Write([]byte("我爱"))
	hasher.Write([]byte("中国"))
	hash := fmt.Sprintf("%x\n", hasher.Sum(nil))
	fmt.Println(hash)

	// sha1/sha256/sha512
	//hasher = sha1.New()
	//hasher.Write([]byte("very good"))
	//hasher.Write([]byte("CHina"))
	//hash = pkg_fmt.Sprint("%x,\n", hasher.Sum(nil))
	//pkg_fmt.Println(hash)
	//
	//pkg_fmt.Print()
}
