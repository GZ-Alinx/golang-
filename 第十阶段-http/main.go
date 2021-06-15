																																																																																																																																																																																package main

																																																																																																																																																																																import (
																																																																																																																																																																																	"encoding/json"
																																																																																																																																																																																	"fmt"
																																																																																																																																																																																	"net/http"
																																																																																																																																																																																	"os"
																																																																																																																																																																																)

																																																																																																																																																																																type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}


func readUsers() []User {
	users := make([]User,0)
	fhandler,err := os.Open(userFile)
	if err != nil {
		return users
	}

	defer fhandler.Close()
	encoder := json.NewDecoder(fhandler)
	err = encoder.Decode(&users)
	if err != nil {
		return users
	}
	return users
}



func main() {

	// url => handler /handlerfunc 先生成html再返回

	// 返回用户列表信息
	/*
	users.json获取用户信息s

	模板 fmt.Sprintf printf Fprintf
	tpl 字符串占位 + 数据 => 结果字符串
	模板技术： 定义模板
		由模板引擎加载模板 通过制定的数据 生成最终的字符串

		模板语法

	go中提供的包(使用方法一样)
		text/template
		html/template  再web开发一定要用这个 会对html中的字符进行转义（防止一些安全注入的问题）

.
	*/
	addr := ":9999"
	http.HandlerFunc("/users/", func(w http.ResponseWriter, request *http.Request) {

	}

	// 启动服务
	http.ListenAndServe(addr, nil)
}
