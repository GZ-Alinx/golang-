package global

import (
	"msg_alert/configdata"

	"go.uber.org/zap"
)

// 全局参数
var (
	Config configdata.ConfigData
	Lg     *zap.Logger
)
