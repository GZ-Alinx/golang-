package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.Set("base.info", "基础配置文件")
	viper.SetConfigName("config1") // 配置文件的文件名 没有扩展名
	viper.SetConfigType("yaml") // 设置扩展名
	viper.AddConfigPath("D:\\golang\\go复习\\标准包学习\\pkg_viper配置解析\\") // 配置文件查找路径
	viper.AddConfigPath(".")
	// 搜索并读取配置文件 之后才能进行使用
	if err := viper.ReadInConfig(); err != nil {
		if _,ok := err.(viper.ConfigFileNotFoundError);ok {
			panic(fmt.Errorf("配置文件没有找到",err))
		} else {
			panic(fmt.Errorf("配置文件找到，但发生了其他错误",err))
		}
	}
	fmt.Println("读取成功")


}
