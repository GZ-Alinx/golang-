package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//"bytes"
)


func getFile(url string) ([]string, error) {

	// 递归拿到目录  递归调用函数 函数自己调用自己
	fileall := []string{}
	files,_ := ioutil.ReadDir(url)
	for _,v := range files {
		//pkg_fmt.Println(i,v)
		if v.IsDir() {
			f,err := getFile(url + "/"+ v.Name())
			if err != nil {
				return nil,err
			}
			fileall = append(fileall, f...)
		}else {
			fileall = append(fileall, url + "/"+ v.Name())
		}
	}

	return fileall,nil
}


func CodeSum(url string) (int,int,error) {
	// 处理文件中代码行数统计
	files,_ := getFile(url)
	//pkg_fmt.Println(files)
	for _,v := range files {
		fmt.Println(v)
		data,  _ := ioutil.ReadFile(v)
		fmt.Println("文件名称:"+string(data))
	}
	count := 0
	fileSum := 0
	for i := 0; i<len(files); i++{
		file, err :=  os.Open(files[i])
		defer file.Close()
		if err != nil {
			fmt.Println("打开文件失败")
			panic(err)
		} else {
			if strings.HasSuffix(files[i], ".windows项目") {
				fileSum += 1
				fd := bufio.NewReader(file)
				for {
					_, err := fd.ReadString('\n')
					//_,err := fd.ReadString('lx')
					if err != nil {
						break
					}
					count++
				}
			}else {
				continue
			}
		}
	}
	//pkg_fmt.Println(count)
	return fileSum, count,nil
}





func main() {
	fmt.Println("代码统计")
	//递归读取特定格式的文件中的换行符号，并对其进行统计
	url := os.Args[1]
	//判断参数个数是否符合规定
	if len(os.Args) != 2 {
		fmt.Println("参数不正确")
		return
	}
	fileSum,codeLines,_ := CodeSum(url)
	fmt.Println("您的go源文件数量为:",fileSum, "您的代码行数为: ",codeLines)

}