package main

import (
	"fmt"
	"strconv"
)

// 单文件管理
var users = []map[string]string{
	{
		"id":   "1",
		"name": "alinx",
		"addr": "深圳",
		"tel":  "19925383400",
	},
	{
		"id":   "2",
		"name": "alinx2",
		"addr": "深圳",
		"tel":  "19925383400",
	},
}

// 获取ID好 并且+1
func GetID() int {
	id := 0
	for _, user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

// 输入函数
func input(prompt string) string {
	fmt.Print(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}

func Add() error {
	// var name, addr, tel string
	// 添加用户 直接赋值
	name := input("请输入用户名:")
	addr := input("请输入地址:")
	tel := input("请输入联系方式:")

	// 验证 用户名不能重复
	if name == "alinx" {
		return fmt.Errorf("用户名%s已存在", name)
	}
	// 切片数据汇总
	users = append(users, map[string]string{
		"id":   strconv.Itoa(GetID()), // 将int转换为字符串
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
	return nil
}

func Query(name string) (string, bool) {
	for _, user := range users {
		if user["name"] == name {
			fmt.Println("用户已存在")
			return user["name"], true
		} else {
			fmt.Println("用户未存在")
			return "not found", false
		}
	}
	return "error", false
}

func Del(name string) bool {
	_, ok := Query(name)
	if ok {
		// 删除用户信息
		fmt.Println("你需要删除的用户是:", name)
		return true
	} else {
		fmt.Println("用户不存在无法删除")
		return false
	}
}

// func Modify(name string) bool {
// 	if Query(name) {

// 	} else {

// 	}
// }

func Action() map[string]string {
	fmt.Println("欢迎使用用户管理系统")
	Cmd := map[string]string{
		"1": "添加用户",
		"2": "删除用户",
		"3": "查询用户",
		"5": "退出用户",
	}
	fmt.Println(Cmd)
	return Cmd
}

func main() {
	Action()
	var a, name, addr, tel string
	fmt.Print("请输入您的操作:")
	fmt.Scan(&a)
	if a == "1" {
		fmt.Printf("添加用户，请输入如下信息")
		fmt.Print("请输入您的name:")
		fmt.Scan(&name)

		fmt.Print("请输入您的addr:")
		fmt.Scan(&addr)

		fmt.Print("请输入您的tel:")
		fmt.Scan(&tel)
		userinfo := map[string]string{
			"id":   "11",
			"name": name,
			"addr": addr,
			"tel":  tel,
		}
		users = append(users, userinfo)
	} else if a == "2" {

	}
}
