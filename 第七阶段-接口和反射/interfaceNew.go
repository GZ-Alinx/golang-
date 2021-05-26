package main

import "fmt"

// 定义结构体
type User struct {
	Id int
	Name string
}


// 定义接口
// 行为： Save保存User切片， Load加载User切片
type Persistent interface {
	Save([]User,string) error
	Load(string)([]User,error)
}



// 结构体
/*
	Gob接口提  定义了Persistent所有的接口，就叫Gob实现了Persistent接口
	此时就能赋值使用了
*/
type Gob struct {

}
// 结构体方法  体现接口不能直接赋值使用


func (p Gob) Save(users []User,path string) error {
	fmt.Println("Gob save")
	return nil
}
func (p Gob) Load(path string) ([]User,error){
	fmt.Println("Gob load")
	return nil,nil
}



// 结构体方法  体现接口不能直接赋值使用 多态测试
type Csv struct {

}


func (p Csv) Save(users []User,path string) error {
	fmt.Println("CSV save")
	return nil
}

func (p Csv) Load(path string) ([]User,error){
	fmt.Println("CSV load")
	return nil,nil
}


func call(persistent Persistent) {
	fmt.Println("call:")
	persistent.Save(nil,"good")
	persistent.Load("good")
}


func main() {
	// 将接口赋值给变量 进行调用
	var persistent Persistent
	fmt.Printf("%T, %#v\n", persistent,persistent)
	//  初始化为 nil

	/* 对于接口来说 不能直接通过接口类型初始化对象
		需要使用实现行为（实现接口的所有方法，某个对象的类型定义了接口中所有的方法）的对象

	*/
	// 赋值使用  Gob结构体实现了 Persistent接口中所有的方法 就能进行对接使用了


	// 多态赋值 得到不同的行为
	persistent =  Csv{} // 赋值给不同的结构体  就能得到不同的行为
	//persistent =  Gob{}
	fmt.Printf("%T, %#v\n",persistent, persistent) // 输出：main.Gob, main.Gob{}
	// 赋值后 persistent就属于Gob结构体的方法对象了  此时就能使用 persistent. 进行调用了

	persistent.Save(nil,"good")
	persistent.Load("good")

	fmt.Println("多态： CSV")
	call(Csv{})
	fmt.Println("多态： Gob")
	call(Gob{})

	// 定义接口实现的对象 是否实现  无语法上的直接关联
	// 只要行为上符合 接口即可关联 这种即是 鸭子类型
	// 可以游泳的 会嘎嘎叫的 就是鸭子
}

