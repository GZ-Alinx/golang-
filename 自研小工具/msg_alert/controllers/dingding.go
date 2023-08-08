package msgs

import (
	"fmt"
	"msg_alert/global"

	"github.com/imroc/req"
	"go.uber.org/zap"
)

func Dingding_Message(message string) string {
	// 设置 Webhook 地址
	webhookUrl := global.Config.WebHook.Dingtalk

	// 使用 req 库向钉钉机器人发送 POST 请求
	_, err := req.Post(webhookUrl, req.BodyJSON(map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": message,
		},
	}))

	if err != nil {
		fmt.Println("发送告警消息失败：", err)
		errs := fmt.Sprintf("Send Faild error: %s", err)
		return errs
	} else {
		fmt.Println("发送告警消息成功！")
		global.Lg.Info("", zap.String("status", "dingtalk send Successfully."))
		return "success"
	}
}
