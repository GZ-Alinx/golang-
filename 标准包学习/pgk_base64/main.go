package main

import (
	"encoding/base64"
	"fmt"	
)

// base64
//     编码/解码 包
//         encoding/base64
func main() {
	// 编码
	ctx := []byte("goodasdfasdfasdf/@5*7asd/+f")
	rs := base64.RawStdEncoding.EncodeToString(ctx)
	fmt.Println(rs)

	// 解码
	txt, _ := base64.RawStdEncoding.DecodeString(rs)
	fmt.Println(string(txt))

}
