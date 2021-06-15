package main

// 读取json配置文件

import (
	"fmt"
	"github.com/spf13/viper"
)

// 定义json结构
type Config struct {
	Redis string
	Mysql Mysqlconfig
}

type Mysqlconfig struct {
	Port int
	Host string
	Username string
	Password string
}



func main() {

	var config Config
	conf := viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("json")
	conf.AddConfigPath(".") // 路径可以设置绝对路径护着当前路径 .

	if err := conf.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	// 读取配置文件信息
	conf.Unmarshal(&config)
	fmt.Println(config.Mysql)
	fmt.Println(config.Redis)



	viper.SetConfigName("config") // 配置文件的文件名 没有扩展名
	viper.SetConfigType("json") // 设置扩展名
	viper.AddConfigPath("D:\\golang\\go复习\\标准包学习\\pkg_viper配置解析\\") // 配置文件查找路径
	viper.AddConfigPath(".")

}

