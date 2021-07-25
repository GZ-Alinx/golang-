package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 请求网络资源
func main() {
	resp,err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	// 请求状态
	fmt.Println(resp.Status)
	fmt.Println(resp.Proto)
	fmt.Println(resp.Header)
	fmt.Println(resp.ContentLength)
	fmt.Println(resp.TLS)
	fmt.Println(resp.ProtoMajor)
	fmt.Println(resp.Cookies())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
