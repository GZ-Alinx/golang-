package main

import (
	"fmt"
	"strings"
)

// 数组
func Typearry() {
	// var 名称[长度]类型
	var names [1]string
	// 输出数据类型
	fmt.Printf("%T\n", names)
	// 实例 使用... 自动推导后面的数据得到数组长度
	name := [...]string{"1", "2", "123", "3", "4", "5"}
	fmt.Printf("%T", name)
	fmt.Println(name)

	// 切片取值
	fmt.Println("切片的值为：", name[2])

	// 对指定位置的数据进行初始化  如下对长度为"3，1，4" 位置的数据进行修改
	passwd := [10]int{3: 1111, 1: 222, 4: 333}
	fmt.Println("访问和修改，长度为3的值：", passwd[3])

	// 获取数组的长度 长度为10
	fmt.Println("数组长度：", len(passwd))

	// 数组切片
	groups := [...]string{"python", "golang", "java", "C#", "groovy"}
	fmt.Println(groups[2:3]) // [开始:结束]

	// 遍历数组
	for i := 0; i < len(groups); i++ {
		fmt.Printf("第[%d]位数: %s\n", i, groups[i])
	}
	// 第二种遍历
	for i, v := range groups {
		fmt.Printf("[%d]: %s\n", i, v)
	}
}

//  多维数组
func Typearrys() {
	// 多为数组的长度只有一维可以自动推导，二维必须定义长度
	langes1 := [...][3]string{{"golang", "python", "groovy"}, {"java", "golang", "C++"}}
	langes2 := [3][3]string{{"windows项目", "shell", "groovy"}, {"java", "golang", "C++"}}
	langes3 := [3][3]string{0: {"windows项目", "shell", "groovy"}, 2: {"java", "golang", "C++"}}
	langes4 := [4][3]string{0: {0: "windows项目", 2: "shell", 1: "groovy"}, 1: {"java", "golang", "C++"}, 3: {"php", "javascriptr", "html"}}

	// 输出类型
	fmt.Printf("数组： \n%T\n%T\n%T\n%T\n", langes1, langes2, langes3, langes4)
	// 输出值
	fmt.Printf("%q\n", langes1)

	// 多为数组访问
	devops := langes1[0][0:2]
	fmt.Printf("运维开发需要掌握的语言：%s\n", devops)
	dev := langes2[1][0:3] // 一般结束位都需要按照下标+1
	fmt.Printf("开发需要掌握的语言：%s\n", dev)

	// 多维数组修改
	fmt.Printf("更正前- 运维开发需要掌握的语言为：%s\n", langes1[0])
	langes1[0][2] = "shell"
	fmt.Printf("更正后- 运维开发需要掌握的语言为：%s\n", langes1[0])

	fmt.Println("----------------------")
	fmt.Printf("更正前- 运维开发需要掌握的技能为：%s\n", langes2[0])
	langes2[0] = [3]string{"linux", "database", "CICD"}
	fmt.Printf("更正后- 运维开发需要掌握的技能为：%s\n", langes2[0])

	// 多为数组遍历
	// langes3 := [3][3]string{0: {"windows项目", "shell", "groovy"}, 2: {"java", "golang", "C++"}}
	for i, line := range langes3 {
		// pkg_fmt.Printf("第一次遍历：%s\n", line)
		for ii, vv := range line {
			// pkg_fmt.Printf("第二次遍历: %s\n", vv)
			fmt.Printf("langes3数组中的值为：[%d %d %s]\n", i, ii, vv) // 此处需要多注意
		}
	}

	// 多为数组遍历第二遍测试
	for i, line := range langes1 {
		// pkg_fmt.Printf("第一次遍历：%s\n", line)
		for ii, vv := range line {
			// pkg_fmt.Printf("第二次遍历: %s\n", vv)
			fmt.Printf("langes1数组中的值为：[%d %d %s]\n", i, ii, vv) // 此处需要多注意
		}
	}
}

// 切片
func Typeslices() {
	// 切片是长度可变的数组，数组在定义好之后长度不可变
	var slice1 []string // 声明切片 nil切片
	fmt.Printf("%T,%t,%v\n", slice1, slice1 == nil, slice1)

	// 初始化 使用make函数初始化切片
	slice2 := make([]string, 10)
	slice3 := []string{"1", "3", "2"}

	arrs := [...][3]string{{"golang", "python", "groovy"}, {"java", "golang", "C++"}}
	fmt.Printf("切片：%q   类型: %T\n", slice3, slice3) // 切片类型
	fmt.Printf("切片：%T， 无固定长度\n", slice2)           // 切片类型
	fmt.Printf("数组：%T， 必须有长度\n", arrs)             // 数组类型
	/* 切片和数据的区别：
	数组是长度固定的元素序列 切片：[]string
	切片是长度可变的元素序列 数组：[2][3]string

	nil 切片，空切片
	slice := []string{"1","2"}
	slice := make([]string{})
	*/

	// 切片操作 初始化  初始化才能用 否则无法使用
	stu := make([]string, 3, 5)
	fmt.Printf("步长: %d\n长度: %d\n", len(stu), cap(stu))
	fmt.Printf("%q,%q\n", stu[0], stu[1:5])
	// 修改
	stu[0] = "golang"
	stu[1] = "python"
	fmt.Printf("%q,%q\n", stu[0], stu[1])
	stu2 := []string{"golang", "java", "python", "C#", "C++"} //  切片定义
	fmt.Printf("%T : %q\n", stu2, stu2)

	// 切片遍历第一种方式
	for index, value := range stu2 {
		fmt.Printf("切片第一种遍历: %d %s\n", index, value)
	}

	// 切片遍历第二种方式
	for i := 0; i < len(stu2); i++ {
		fmt.Printf("切片第二种遍历: %s\n", stu2[i])
	}

	// 切片添加元素 append
	stu2 = append(stu2, "shell")
	fmt.Println(stu2)
	// 切片删除元素
	// stu2[start:end] stu2中，从start开始到end-1所有元素组成的切片
	stu2 = stu2[0:4] // 以索引进行删除
	fmt.Println(stu2)
	// 删除第一个元素，删除最后一个元素
	stu3 := []string{"golang", "java", "python", "C#", "C++"} //  切片定义
	stu4 := stu3[1:len(stu3)]
	fmt.Println("-----------------")
	fmt.Println(stu3)
	fmt.Printf("stu4: %s\n", stu4)
	stu5 := stu3[0:len(stu3)]
	fmt.Printf("stu5: %s\n", stu5)

	// 删除中间一个数据 删除3
	nums1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// copy(dst,src) src > dst
	nums2 := []int{10, 11, 12, 13, 15, 15, 16}
	copy(nums1, nums2) // 按照索引进行复制对应起来
	fmt.Println(nums1, nums2)

	// 切片追加

	//# 可以将string拼接到byte切片
	slice := make([]byte, 10)
	var a string = "Hi "
	var b string = "YY"
	slice = append([]byte(a), b...)
	fmt.Println("切片追加，解包的方式： ", string(slice))
}

// 映射map 类似py字典
func Typemaps() {
	// 定义map 第一种
	var Iaas map[string]string                                           // 声明
	Iaas = map[string]string{"name": "alinx", "age": "28", "site": "cn"} // 定义
	// Iaas["name"] = "888"
	// 直接使用make进行初始化创建map
	// 定义map 第二种
	man := make(map[string]string)
	fmt.Printf("%T,%q\n%T\n", Iaas, Iaas, man)
	// 定义map 第三种
	wom := map[string]string{}
	fmt.Printf("%T\n", wom)

	// 操作map
	langes := map[string]string{"dev": "golang,java,,,,python", "ops": "shell,python,golang", "devops": "golang,shell"}
	fmt.Println(strings.Split(langes["dev"], ",")) // 数据分割
	fmt.Printf("长度: %d\n", len(langes))

	// map 取值
	dev, ok := langes["dev"]
	fmt.Printf("map取值：\n  存在吗:%t\n  取出值:%v\n", ok, dev)

	// 修改
	langes["ops"] = "java,C++"
	fmt.Printf("修改后的值：%s\n", langes)

	// 删除
	delete(langes, "dev")
	fmt.Printf("删除后的值：%s\n", langes)

	// 遍历map(映射)
	for k, v := range langes {
		fmt.Printf("遍历map数据: %s %s\n", k, v)
	}


	// map 追加

}

func main() {
	Typearry()
	Typearrys()
	Typeslices()
	Typemaps()
}
