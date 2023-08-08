package configdata

type ConfigData struct {
	LogfilePath string `mapstructure:"logfile_path"`
	WebHook     Hook   `mapstructure:"Webhook"`
}
type Hook struct {
	Wechat   string `mapstructure:"Wechat"`
	Feishu   string `mapstructure:"Feishu"`
	Dingtalk string `mapstructure:"Dingtalk"`
}
