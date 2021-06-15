package main


import (
	_ "cmdb/routers"
	"net/http"
)

func main() {
	// 用户登录逻辑


	// url => handler/handlerfunc => template加载模板并输出
	// 定义Handler/HandlerFunc
	// 绑定关系 url handler
	// 启动服务
	addr := ":9999"
	http.ListenAndServe(addr, nil)


}
