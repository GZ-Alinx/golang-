package  main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



type ResGroup struct {
	Data string
	Path string
}

type Info struct {
	Data string
	Path string
}


func main() {
	// 初始化示例
	route := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	// 路由分组格式和写法
	v1 := route.Group("/v1")
	{
		v1.GET("/login1", login1)
		v1.GET("/r2", info1)
	}

	v2 := route.Group("/v2")
	{
		v2.GET("/login2", login2)
		v2.GET("/r4", info2)
	}
}


func login1(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"login1", c.Request.URL.Path})
}

func info1(c *gin.Context) {
	c.JSON(http.StatusOK,Info{"info1", c.Request.URL.Path})
}



func login2(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"login2", c.Request.URL.Path})
}

func info2(c *gin.Context) {
	c.JSON(http.StatusOK, Info{"info2", c.Request.URL.Path})
}