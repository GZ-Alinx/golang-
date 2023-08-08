package main

import (
	"fmt"
	msgs "msg_alert/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 消息接收结构体
type MessageType struct {
	Type    string `json:"type"`
	Context string `json:"context"`
}

func MSG_SEND(ctx *gin.Context) {
	var msg MessageType
	if err := ctx.ShouldBind(&msg); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "数据绑定失败 - Send Failure",
		})
	} else {
		fmt.Println("请求信息： ", msg)
		if msg.Type == "Feishu" {
			response := msgs.Feishu_Message(msg.Context)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": response,
			})
		} else if msg.Type == "Weixin" {
			response := msgs.WeChat_Message(msg.Context)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": response,
			})
		} else if msg.Type == "Dingtalk" {
			msgs.Dingding_Message(msg.Context)

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":    "不支持的消息类型",
				"result": "支持的类型: Feishu, Weixin, Dingtalk ",
			})
		}
	}
}

func main() {
	fmt.Println("消息发送服务")
	r := gin.Default()
	r.POST("/send", MSG_SEND)
	r.Run(":9090")
}
