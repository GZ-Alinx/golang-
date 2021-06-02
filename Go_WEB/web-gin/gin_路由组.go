package main


import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 路由组
	v1 := router.Group("/v1")
	{
		v1.GET("login", loginEndpoint)
		v1.POST("submit", submitEndpoint)
		v1.POST("/read",readEndpoint)
	}

	v2 := router.Group()
	v2.GET("login", loginEndpoint)
	v2.POST("submit", submitEndpoint)
	v2.POST("/read",readEndpoint)

	router.Run(":8080")

}