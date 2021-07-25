package main


// gin 框架的bind是将请求的数据赋值到结构体中生成可调用的对象(类似于序列化和反序列化)

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 将绑定的数据存储在结构体中
type Login struct {
	// 绑定 binding:"required" 不能为空
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remark string `json:"remark"`
}
func main() {
	r := gin.Default()
	r.POST("/login", func(context *gin.Context) {
		var login Login
		err := context.Bind(&login) // 将结构体数据绑定到对象中
		fmt.Println("请求数据:",login)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg":"绑定失败,参数错误",
				"data":err.Error(),
			})
			return
		}
		// 如果请求的数据正常 将进行判断
		if login.UserName == "alinx" && login.Password=="1234" {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg":"登录成功",
				"data":"OK",
			})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":"登录失败",
			"data":"error",
		})
		return
	})

	r.Run(":9090")
}


/*请求方式
POST 127.0.0.1:9090/login

请求参数
{
    "user_name":"alinx",
    "password":"1234",
    "remark":"测试请求"
}

*/