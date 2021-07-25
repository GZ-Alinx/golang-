package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cmdb?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("链接异常")
	}
	err := db.ping()
	if err != nil {
		fmt.Println("连接异常")
	}
	//return db
}