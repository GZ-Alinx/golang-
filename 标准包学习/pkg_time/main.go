package main

import (
	"fmt"
	"time"
)

func main() {
	// time
	// 时间: 字符串格式 unix时间戳
	now := time.Now()
	fmt.Printf("%T,\n %#v", now, now)

	fmt.Println() // 时间戳
	timestarup := time.Now().Unix()
	tm := time.Unix(timestarup, 0)
	times := tm.Format("2006-01-02 15:03:04")
	fmt.Println(times)
	// 时间格式化
	fmt.Println(now.Format("2006-01-02 03:04:05")) // 12进制
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 24进制

	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
}
