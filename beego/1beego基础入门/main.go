package main

import (
	"github.com/astaxie/beego"
)

func main() {
	/*
		beego官方网站: https://beego.me/
			目前版本
				2.0.0 bate
				1.12.0

			使用
			1. 定义处理器
			2. 绑定URL和处理器的功能 简称路由
				基本路由
					为url帮i的那个某个请求方法的路由参数
					为根路劲的GET方法绑定函数
					beego.Get("/url", 绑定函数)
			3. 启动服务
	*/
	beego.Run()
	//访问http://127.0.0.1:8080/

}
