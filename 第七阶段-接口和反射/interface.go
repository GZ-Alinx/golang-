package main

import "fmt"

// 接口定义

type User struct {
	Id int
	Name string
}

// 定义接口Persistent
// 行为： Save保存User切片
// 方法名称  返回值类型
type Persistent interface {
	Save([]User,string) error
	Load(string) ([]User,error)
	// 接口中所定义的方法，需要有确切的实现，否则异常
}

type GObPersustent struct {}
// 定义了GObPersustent 所有接口 就叫GObPersustent实现了Persistent接口

func (p GObPersustent) Save(users []User,path string) error {
	fmt.Printf("\ngob persistent save")
	return nil
}


func (p GObPersustent) Load(path string) ([]User,error) {
	fmt.Printf("\ngob persistent load")
	return nil,nil
}

// 作用
type CsvPersistent struct{}

func (p CsvPersistent) Save(users []User,path string) error {
	fmt.Println("csv persistent save")
	return nil
}

func (p CsvPersistent) Load(path string) ([]User, error) {
	fmt.Println("csv persistent load")
	return nil,nil
}


// 此处使用接口作为类型  对接起来即可
func call(persistent Persistent) {
	fmt.Println("\ncall:")
	persistent.Save(nil,"")
	persistent.Load(":")
}



func main() {
	var persistent Persistent

	fmt.Printf("%T,%#v\n", persistent, persistent)

	// 赋值
	// 接口 不能直接通过接口类型进行初始化对象
	// 需要使用实现行为（实现接口的所有方法，某个对象的类型定义了接口中所有的方法）的对象
	persistent = GObPersustent{}
	//persistent = CsvPersistent{}

	fmt.Printf("%T, %#v\n",persistent,persistent)
	persistent.Save(nil,"")
	persistent.Load("")
	call(GObPersustent{})
	call(CsvPersistent{})
	// 定义类型 与接口是否实现  无语法上直接关联
	//鸭子类型： 你怎么判断一个动物是否是一个鸭子
	//	根据行为判断： 游泳，嘎嘎叫 都是鸭子

	// 根据以上知识  只要实现了接口中定义的方法 那就是接口类型






	// 指针对象
	persistent = new(CsvPersistent)
	fmt.Println("指针类型：\n\n%T,%#v\n", persistent,persistent)
	persistent.Load("")
	persistent.Save(nil,"")
	call(&CsvPersistent{})
	call(&GObPersustent{})


	// 在类型上 值类型会自动 生成指针类型  指针接收者无法生成值接收者

}

