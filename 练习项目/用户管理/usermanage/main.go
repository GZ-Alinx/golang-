package main

import (
	"fmt"
	"usermanage/add"
	"usermanage/del"
	"usermanage/modify"
	"usermanage/query"
)

var (
	name    string
	sex     string
	phone   int
	address string
)

user := make([]string{})

func main() {
	fmt.Println("环境使用用户管理系统")
	add.Add()
	del.Del()
	query.Query()
	modify.Modify()
}
