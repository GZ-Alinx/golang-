package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// gin基础
func main() {
	// 初始化gin
	router := gin.Default()
	//router.Any("/", WebRoot) // 设置路由
	//router.GET("/test",WebRoot)

	//func WebRoot(contexts *gin.Context) {
	//	contexts.String(http.StatusOK, "hello world")
	//}

	// 注册一个动态路由
	router.GET("/user/:name", func(c *gin.Context) {
		// 使用c.Param(key) 获取url参数
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s", name)
	})

	// 高级动态路由
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK, message)
	})

	// 启动并绑定端口 默认端口8080
	router.Run(":8080")
}

