package main

import (
	"fmt"
	"io"
	"net/http"
)


// 客户端
func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("请求异常")
	}
	fmt.Println(resp.Body.Close())
	body,err := io.ReadAll(resp.Body)
	fmt.Println(body)
}
