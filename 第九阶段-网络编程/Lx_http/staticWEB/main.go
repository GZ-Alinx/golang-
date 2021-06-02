package main

import "net/http"


// web服务端 静态资源创建
func main() {
	addr := ":9999"
	// URL 映射目录
	//http.Handle("/static1/", http.FileServer(http.Dir("./www")))
	//http.Handle("/static2/", http.FileServer(http.Dir("./www")))
	//http.Handle("/static3/", http.FileServer(http.Dir("./www")))

	//简写为此
	// Handler url处理的 主要函数
	// HandlerFunc


	http.Handle("/static/",
		http.StripPrefix(
			"/static",
			http.FileServer(http.Dir("./www")),
			),
		)
	http.ListenAndServe(addr, nil)
	//www 是所有资源的路径
	//static1 是具体的目录路径
	//目录结构 www/static1



}