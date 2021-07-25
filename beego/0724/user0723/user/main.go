package main

import (
	"log"

	_ "user/routers"
	"user/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := beego.AppConfig.String("mysql::DSN") // 通过mysql的配置读取
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err) // 记录错误并退出
		// default 数据库链接别名  多个数据建库链接
	}

	// 日志配置
	/*
		主要的参数如下说明：
			filename 保存的文件名
			maxlines 每个文件保存的最大行数，默认值 1000000
			maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
			daily 是否按照每天 logrotate，默认是 true
			maxdays 文件最多保存多少天，默认保存 7 天
			rotate 是否开启 logrotate，默认是 true
			level 日志保存的时候的级别，默认是 Trace 级别
			perm 日志文件权限
	*/
	// web.BConfig.Log.AccessLogs = false
	logs.SetLogger(logs.AdapterFile,
		`{
		"filename":"logs/web.log",
		"level":7,
		"maxlines":0,
		"maxsize":0,
		"daily":true,
		"maxdays":365,
		"color":false,
		}`,
	)

	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")

	// orm配置
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	if user := services.GetUserByName("admin"); user == nil {
		log.Print("创建管理员账号")
		services.AddUser("admin", "123@abc", "深圳", true)
	}

	addr := ":8888"
	// beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.56.103:6379"

	beego.Run(addr)
}
