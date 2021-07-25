package main

import (
	_ "beego0718/routers"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	dsn, err := web.AppConfig.String("mysql::dsn")
	if err != nil || dsn == "" {
		log.Fatal("error config dsn")
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	web.Run()
}
