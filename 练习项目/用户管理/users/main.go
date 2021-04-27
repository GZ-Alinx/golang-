package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 切片定义，值为map类型
var users = []map[string]string{}

var id int = 0

func init() {
	for i := 1; i < 10; i++ {
		i := strconv.Itoa(i) // 将int转换为字符串
		// strconv.Atoi(i) // 将字符串转换为int
		fmt.Printf("%T,%#v\n", i, i)
		users = append(users, map[string]string{"id": "1" + i, "name": "alinx" + i, "add": "CN" + i, "tel": "138000" + i})
	}
}

func Get_id() string {
	for _, v := range users {
		if strconv.Atoi(v["id"]) > id {
			id = strconv.Atoi(v["id"])
		}
	}
	fmt.Printf("%T,ID:", id, id)
	return id
}

func add() map[string]string {
	var (
		// id   string
		name string
		add  string
		tel  string
	)
	fmt.Print("user:")
	fmt.Scan(&name)

	fmt.Print("add:")
	fmt.Scan(&add)

	fmt.Print("tel:")
	fmt.Scan(&tel)
	user := map[string]string{
		"id":   string(Get_id()),
		"name": name,
		"add":  add,
		"tel":  tel,
	}
	return user

}

func Query(names string) string {
	for _, v := range users {
		// fmt.Printf("%T\n,%#v\n", v, v)
		if strings.Contains(v["name"], names) { // 数据包含
			fmt.Print("数据已找到: ")
			fmt.Println(v)
			return "数据已找到"
		}
	}
	return "不包含"
}

func Delete(name string) bool {
	for k, v := range users {
		fmt.Print(k, v)
	}
	return true
}

const password = "123"

func hash(passwd string) bool {
	ctx := []byte(passwd)
	hash := fmt.Sprintf("%x\n", md5.Sum(ctx))
	ctx2 := []byte(password)
	hash2 := fmt.Sprintf("%x\n", md5.Sum(ctx2))
	fmt.Println(hash, hash2)

	sha1s := sha1.Sum([]byte(passwd))
	fmt.Println("sha1 算法:", sha1s)
	if hash != hash2 {
		fmt.Println("密码错误")
		return false
	}
	fmt.Println("密码正确")
	fmt.Println(Get_id())
	return true
}

func main() {
	var Password string
	fmt.Printf("输入密码:")
	fmt.Scan(&Password)
	if !hash(Password) {
		fmt.Println("密码错误，请输入正确的密码!")
		os.Exit(1)
	}

	action := ""
	for {
		fmt.Print("欢迎使用xxx系统\n1. 添加数据\n2. 查询数据\n0. 退出系统\n")
		fmt.Print("输入你的操作:")
		fmt.Scan(&action)
		if action == "1" {
			fmt.Println("添加数据")
			users = append(users, add())
		} else if action == "2" {
			fmt.Println("查询数据")
			name := ""
			fmt.Print("输入查询名字:")
			fmt.Scan(&name)
			Query(name)
		} else if action == "0" {
			fmt.Println("退出")
			os.Exit(0)
		} else {
			fmt.Println("操作有误")
		}
		// fmt.Println(users)
	}
}
