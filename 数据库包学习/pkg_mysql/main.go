package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "88888888"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "django"
)

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	// 下载安装
	//  go get "github.com/go-sql-driver/mysql"

	/*
		特性
			轻巧快速
			本机Go实施。没有C绑定，只有纯Go
			通过TCP / IPv4，TCP / IPv6，Unix域套接字或自定义协议的连接
			自动处理断开的连接
			自动连接池（按数据库/ sql程序包）
			支持大于16MB的查询
			全力sql.RawBytes支持。
			LONG DATA在准备好的语句中进行智能处理
			LOAD DATA LOCAL INFILE通过文件允许列表和io.Reader支持获得安全支持
			可选time.Time解析
			可选的占位符插值

		环境要求
			达到1.10或更高。我们旨在支持Go的3个最新版本。
			MySQL（4.1 +），MariaDB，Percona Server，Google CloudSQL或Sphinx（2.2.3+）
	*/

	// 使用
	// "database/sql"

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	MysqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	MysqlDb.SetConnMaxIdleTime(100 * time.Second) // 连接最大周期，超过即超时
	MysqlDb.SetMaxOpenConns(1000)                 // 设置连接数
	MysqlDb.SetMaxIdleConns(16)                   // 设置限制连接数

	fmt.Println(MysqlDb.Ping())

	// 查询数据，指定字段名

	user := new(User)
	row := MysqlDb.QueryRow("select id, name, age from users where id=?", 1)
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(user.Id, user.Name, user.Age)
}
