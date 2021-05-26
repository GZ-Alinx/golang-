package main

import "fmt"

// 用户管理系统
type UserManager interface {
	Add(user Users) ([]Users, error)
	Del(user string) (string,error)
	Change(u string, user Users) (Users,error)
	Query(user string) (Users,error)
}

type Users struct {
	Id int
	Name string
	Age int
	Work string
	Site string
}

var USERS = []Users{}

func (users Users) Add(user Users) ([]Users,error){
	USERS = append(USERS, Users{user.Id,user.Name,user.Age,user.Work,user.Site})
	return USERS,nil
}

func (users Users) Del(user string) (string, error) {
	for k,v := range USERS {
		fmt.Printf("%#v, %#v\n",k,v)
		if v.Name == user {
			fmt.Println("用户存在，准备删除")
			USERS = append(USERS[:k], USERS[k+1:]...)
		}
	}
	return user,nil
}

func (users Users) Change(u string, user Users) (Users, error) {
	for k,v := range USERS {
		if v.Name == u {
			fmt.Println("用户存在，继续变更中")
			USERS = append(USERS[:k], USERS[k+1:]...)
			data := Users{user.Id,user.Name,user.Age,user.Work,user.Site}
			USERS = append(USERS,data)
			break
		}else {
			fmt.Println("用户不存在无法变更")
			break
		}
	}
	return user,nil
}


func (users Users) Query(user string) (Users, error) {
	for _,v := range USERS {
		if v.Name == user {
			return Users{v.Id,v.Name,v.Age,v.Work,v.Site}, nil
		}else {
			fmt.Println("你查询的用户不存在")
			return users,nil
		}
	}
	return users,nil
}




func main() {
	// Users结构体实现了 UserManager接口中定义的四种方法 所以可以这样进行定义使用
	var userManager UserManager

	// 两种方法进行
	userManager = Users{} // 值类型
	//userManager = new(Users) //指针类型

	userManager.Add(Users{1,"alinx11",27,"SCM","CN"})
	userManager.Add(Users{1,"alinx22",27,"SCM","CN-GZ"})
	userManager.Add(Users{1,"alinx24",27,"SCM","CN-GZ"})
	userManager.Add(Users{1,"alinx25",27,"SCM","CN-GZ"})
	userManager.Add(Users{1,"alinx26",27,"SCM","CN-GZ"})
	userManager.Del("alinx11")
	userManager.Change("alinx22",Users{1,"改名卡",2801,"SCMsss","CNs"})
	userManager.Query("lixiong")
	fmt.Println(USERS)

}