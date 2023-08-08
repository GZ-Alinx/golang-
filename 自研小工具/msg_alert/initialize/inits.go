package initialize

import (
	"fmt"
	"time"

	"msg_alert/global"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

// InitLogger
func InitLogger() {
	// zap instance
	cfg := zap.NewDevelopmentConfig()

	// log stdout path and file name set.
	cfg.OutputPaths = []string{
		fmt.Sprintf("%smsg_alert-%s.log", global.Config.LogfilePath, GetNowFormatTodayTime()),
		"stdout",
	}

	// create logger instance
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
	color.Green("log config init Successfully.")
}

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}
