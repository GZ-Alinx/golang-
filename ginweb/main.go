package  main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化示例
	r := gin.Default()

	// 对接外部处理函数
	r.GET("/health", HealthFunc)

	// 对接内部处理函数
	r.POST("/POST", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"msg":"响应成功",
		})
	})
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
