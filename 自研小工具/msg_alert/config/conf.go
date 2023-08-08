package config

import (
	"msg_alert/configdata"
	"msg_alert/global"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// config init
func InitConfig() {
	// Use viper init configure
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		color.Red("read file failed， please check if the configuration file is correct")
		panic(err)
	}
	serverConfig := configdata.ConfigData{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		color.Red("format failed!!， please check if teh configuration file is correct")
		panic(err)
	}

	global.Config = serverConfig
	color.Green("Configuration init Successfully.")
}
