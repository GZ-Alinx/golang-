package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context" // 获取请求数据 都需要通过此方法
)

func main() {
	// 为URL绑定某个请求方法的路由函数
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hi, beego"))
	})
	/* 请求方式
	Post，Delete，Head，PUT，OPtions，Patch

	curl 请求: curl -XPOST "http://localhost:8080" -v
	*/
	beego.Post("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("post"))
	})

	beego.Delete("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Delete"))
	})

	// 绑定所有请求方法 如果没有定义优先级将使用any绑定的方法进行请求
	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("any"))
	})

	/*
		提交参数
		常规:  ?a=b
		正则路由: /delete/1/  可使用正则
		注意:  使用`` 进行引用编写正则才能生效
	*/

	// \d{1,8} 数字 1-8位
	beego.Any(`/delete/:id(\d{1,8})/`, func(ctx *context.Context) {
		// 获取参数  ctx.Input.Param(":id")
		ctx.Output.Body([]byte("delete args" + ctx.Input.Param(":id")))
		// http://127.0.0.1:8080/delete/123 请求地址
	})
	// 另外一种写法
	beego.Any(`/get/:id:int/`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("delete args" + ctx.Input.Param(":id")))
	})

	// 启动服务
	beego.Run()
}
