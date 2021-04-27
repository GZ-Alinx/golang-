package main

import (
	"fmt"
	"net/http"
)

// 面向过程 处理器函数
// 面向对象 处理器

// 处理器函数
func Home(w http.ResponseWriter, r *http.Request) {
	// 用户提交的数据 http内容  go代码转换 http.Request
	w.Write([]byte("hi"))
}

// 自定义类型 处理器
type Help struct{}

func (h *Help) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "help")
}

// http 提供 http客户端和服务端的实现

func main() {
	addr := ":8888"

	// 绑定
	http.HandleFunc("/home", Home) // 处理器函数绑定
	// http.Handle("/help", new(Help)) // 处理器绑定
	http.ListenAndServe(addr, nil)
	// http.Dir 类型转换
	// http.Handle("/w1", http.FileServer(http.Dir("./static")))
	// http.Handle("/www", http.StripPrefix("/www", http.FileServer(http.Dir("./static"))))
}
