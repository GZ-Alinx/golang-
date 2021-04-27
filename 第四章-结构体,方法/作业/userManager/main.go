package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for _, v := range users {
		if name == v["name"] {
			fmt.Errorf("用户%s已存在", name)
			return fmt.Errorf("用户%s已存在", name)
		}
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

func Query(name string) map[string]string {
	for _, user := range users {
		if user["name"] == name {
			fmt.Println("查询成功~")
			fmt.Println(user)
			return user
		}
	}
	fmt.Println("查询失败")
	return map[string]string{}
}

func Del(name string) bool {
	// 删除用户信息
	for i, v := range users {
		if v["name"] == name {
			fmt.Println("你需要删除的用户是:", name)
			users = append(users[:i], users[i+1:]...)
			fmt.Println("删除成功")
			return true
		} else {
			fmt.Println("用户不存在无法删除")
			return false
		}
	}
	return false
}

func Action() map[string]string {
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println("欢迎使用用户管理系统")
	fmt.Println(strings.Repeat("*", 20))

	Cmd := map[string]string{
		"1": "添加用户",
		"2": "删除用户",
		"3": "查询用户",
		"5": "退出用户",
	}

	for k, v := range Cmd {
		fmt.Println(k, v)
	}
	// 排序sort
	//less := func(i,j string) bool {
	//	return Cmd[i] < Cmd[j]
	//}
	//fmt.Println(Cmd, less)
	//fmt.Println(sort.Slice(Cmd, less))
	return Cmd
}

func main() {
	var a, name, addr, tel string
	for {
		Action()
		fmt.Print("请输入您的操作:")
		fmt.Scan(&a)
		if a == "1" {
			fmt.Printf("您选择的是 添加用户 \n请输入如下信息")
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
			fmt.Println(users)
			continue
		} else if a == "2" {
			fmt.Println("您选择的是 删除用户")
			fmt.Print("请输入删除的name:")
			fmt.Scan(&name)
			Del(name)
			fmt.Println(users)
			continue
		} else if a == "3" {
			fmt.Println("您选择的是 查询用户")
			fmt.Print("请输入查询的用户名称:")
			fmt.Scan(&name)
			Query(name)
			continue
		} else if a == "4" {

			continue
		} else if a == "5" {
			fmt.Println("退出系统")
			os.Exit(0)
		} else {
			fmt.Println("输入不正确，请重新输入")
		}
	}
}
