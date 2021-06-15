package main



import (
	"github.com/gin-gonic/gin"

)

// 请求路由  控制ui的资源获取
func main() {
	//请求路由绑定静态文件夹，参数作为url 泛绑定
	r := gin.Default()
	

	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "post")
	})

	r.Handle("DELETE", "/deletes", func(c *gin.Context) {
		c.String(200, "delete")
	})

	// 多种请求类型 对应一个处理方法
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run()
}
