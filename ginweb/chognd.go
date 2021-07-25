	package  main

	import (
		"github.com/gin-gonic/gin"
		"net/http"
	)

	func main() {
		// 初始化示例
		r := gin.Default()


		// 重定向到百度
		r.POST("/rewrite1", func(c *gin.Context) {
			url := "https://www.baidu.com"
			c.Redirect(http.StatusMovedPermanently, url)
		})

		// 重定向到 health 路由
		r.POST("/rewrite2", func(c *gin.Context) {
			c.Request.URL.Path = "/health"
			r.HandleContext(c)
		})

		r.POST("/health", HealthFunc)



		// 默认端口是 8080
		r.Run(":9090")
	}


	// 处理函数
	func HealthFunc(c *gin.Context) {
		// 返回JSON格式数据
		c.JSON(http.StatusOK, gin.H{
			"message":"Running",
		})
	}
