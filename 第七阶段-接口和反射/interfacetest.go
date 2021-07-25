package main

import "fmt"

// 接口定义
type User struct {
	Id int
	Name string
}

// 定义接口 Persistent

/* 行为：
Save保存
User切片
*/


type Persistent interface {
	// 方法名称  返回值类型  接口中所定义的方法，需要有确切的实现，否则异常
	Save([]User,string) error
	Load(string) ([]User,error)
}


// 实现原则： 定义了GObPersustent 所有接口 就叫GObPersustent实现了Persistent接口



// 第一个结构体实现 Persistent接口
type GObPersustent struct {}


// 基于GObPersustent结构体 实现Save方法
func (p GObPersustent) Save(users []User,path string) error {
	fmt.Printf("\ngob persistent save")
	return nil
}
// 基于GObPersustent结构体 实现Load方法
func (p GObPersustent) Load(path string) ([]User,error) {
	fmt.Printf("\ngob persistent load")
	return nil,nil
}




// 第二个结构体实现 Persistent接口
type CsvPersistent struct{}

// 基于CsvPersistent结构体 实现Save方法
func (p CsvPersistent) Save(users []User,path string) error {
	fmt.Println("csv persistent save")
	return nil
}
// 基于CsvPersistent结构体 实现Load方法
func (p CsvPersistent) Load(path string) ([]User, error) {
	fmt.Println("csv persistent load")
	return nil,nil
}




// 此处使用接口作为类型  对接起来调用测试
func call(persistent Persistent) {
	fmt.Println("\ncall:")
	persistent.Save(nil,"")
	persistent.Load(":")
}



// 调用CsvPersistent结构体实现了 Persistent 接口的方法
func Csv() {
	var persistent Persistent
	fmt.Printf("%T,%#v\n", persistent, persistent)
	persistent = CsvPersistent{}
	persistent.Load("")
	persistent.Save(nil,"")
	/*
	   接口 不能直接通过接口类型进行初始化对象
	   结构体使用实现行为（实现接口的所有方法，某个对象的类型定义了接口中所有的方法）的对象
	*/
}

// 调用GObPersustent结构体实现了 Persistent 接口的方法
func GOb() {
	var persistent Persistent
	fmt.Printf("%T,%#v\n", persistent, persistent)
	persistent = GObPersustent{}
	persistent.Load("")
	persistent.Save(nil,"")
	/*
	   接口 不能直接通过接口类型进行初始化对象
	   结构体使用实现行为（实现接口的所有方法，某个对象的类型定义了接口中所有的方法）的对象
	*/
}




func main() {
	// 调用  默认是值类型调用
	Csv()
	GOb()


	/*
	   接口 不能直接通过接口类型进行初始化对象
	*/


	// 在类型上 值类型会自动 生成指针类型  指针接收者无法生成值接收者
	var persistent Persistent
	fmt.Printf("%T,%#v\n", persistent, persistent)


	// 函数中调用使用
	call(GObPersustent{})
	call(CsvPersistent{})
	/*
	    定义类型与接口是否实现  无语法上直接关联
			鸭子类型： 你怎么判断一个动物是否是一个鸭子
			根据行为判断： 游泳，嘎嘎叫 都是鸭子

		总结： 根据以上知识  只要结构体实现了接口中定义的方法 那就是接口类型，就能赋值进行调用
	*/


	// 指针对象  自动生成的指针类型调用
	persistent = new(CsvPersistent)
	fmt.Printf("指针类型：\n\n%T,%#v\n", persistent,persistent)
	persistent.Load("")
	persistent.Save(nil,"")
	call(&CsvPersistent{})
	call(&GObPersustent{})

}
