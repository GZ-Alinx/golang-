package msgs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"msg_alert/global"
	"net/http"
)

func Feishu_Message(msg string) string {
	// 设置 Webhook 地址
	webhookUrl := global.Config.WebHook.Feishu

	// 要发送的消息内容
	message := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": msg,
		},
	}
	// 转换消息内容为JSON格式
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "failed"
	}

	// 发送HTTP POST请求
	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "failed"
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		return "failed"
	}

	fmt.Println("Message sent successfully!")
	return "success"
}
