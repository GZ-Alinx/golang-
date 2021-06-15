package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"encoding/hex"
)




func md5sum(files string) string {
	f, err := os.Open(files)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	md5data := hex.EncodeToString(h.Sum(nil))
	return md5data

}

// 文件获取接口设计

type Fileinfo interface {
	Source(srcpath string ) ([]map[string]string,error)
	Destination(destpath string ) ([]map[string]string,error)
}

type Infos struct {}

func (infos Infos) Source(srcpath string ) ([]map[string]string,error) {
	SOURCE_PATH := make([]map[string]string, 0)
	filepath.Walk(srcpath, func(paths string, info fs.FileInfo, err error) error {
		var Isdir string
		if !info.IsDir() {
			Isdir = "false"
		}else {
			Isdir = "true"
		}
		data := map[string]string{"Name":info.Name(),"IsDir":Isdir,"Path":paths}
		SOURCE_PATH = append(SOURCE_PATH, data)
		return nil
	})
	return SOURCE_PATH,nil
}


func (infos Infos) Destination(destpath string ) ([]map[string]string,error) {
	DEST_PATH := make([]map[string]string, 0)
	filepath.Walk(destpath, func(paths string, info fs.FileInfo, err error) error {
		var Isdir string
		if !info.IsDir() {
			Isdir = "false"
		}else {
			Isdir = "true"
		}
		data := map[string]string{"Name":info.Name(),"IsDir":Isdir,"Path":paths}
		// 此处使用字典嵌套切片的方式进行获取
		//fmt.Println(data)
		DEST_PATH = append(DEST_PATH, data)
		return nil
	})
	return DEST_PATH,nil
}


// 文件校验接口设计
type FileDiff interface {
	Md5Sum(srcpath string, destpath string) ([]string,error)
	BaseSum(srcpath string, destpath string) ([]string,error)
}

type Diff struct {}

// md5计算对比
func (diff Diff) Md5Sum(srcpath string, destpath string) ([]string, error) {
	var Getinfo Fileinfo
	Getinfo = Infos{}
	srcd,_ := Getinfo.Source(srcpath)
	destd,_ := Getinfo.Destination(destpath)
	Differr := make([]string, 0)
	// 提取文件 进行md5值计算
	for _,srcinfo := range srcd {
		for _,destinfo := range destd {
			if srcinfo["IsDir"] == "false" && srcinfo["Name"] == destinfo["Name"] {
				fmt.Println(srcinfo["Path"], destinfo["Path"])
				srcmd5 := md5sum(srcinfo["Path"])
				destmd5 := md5sum(destinfo["Path"])
				//fmt.Println("升级包MD5",srcmd5, "应用内MD5",destmd5)
				if srcmd5 != destmd5 {
					data := "[ 应用文件:"+ destinfo["Path"] + "  升级包文件:" + srcinfo["Path"] + " ]"
					Differr = append(Differr, data)
					//fmt.Printf("应用文件[ %s ] MD5值: %s\n升级包文件[ %s ] MD5值: %s 对比MD5值异常\n", destinfo["Path"],destmd5, srcinfo["Path"], srcmd5)
				}else {
					fmt.Printf("应用文件[ %s ] MD5值: %s\n升级包文件[ %s ] MD5值: %s 对比MD5值正常\n", destinfo["Path"],destmd5, srcinfo["Path"], srcmd5)
				}
			}
		}
	}
	return Differr,nil
}

// base64值对比
func (diff Diff) BaseSum(srcpath string, destpath string ) ([]string, error) {
	return nil,nil
}

// 单文件对比
func fileDiffs(file1, file2 string) []string {
	f1 := md5sum(file1)
	f2 := md5sum(file2)
	if f1 == f2 {
		fmt.Println("文件对比正常")
	}else {
		fmt.Println("文件对比异常")
	}
	data := []string{f1,f2}
	return data
}



func main() {
	// 参数检查
	if len(os.Args) == 1 && len(os.Args) <= 2 {
		fmt.Println("参数不齐全，请加上升级包路径和应用路径")
		os.Exit(0)
	}
	if len(os.Args) >= 4 {
		fmt.Println("参数不齐全，请加上升级包路径和应用路径")
		os.Exit(0)
	}

	// 使用帮助
	if os.Args[1] == "-h"  || os.Args[1] == "-H" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" || len(os.Args) == 1{
		fmt.Println("使用示例:\n    Md5Diff /opt/uptade/ /home/app/crm/")
		os.Exit(0)
	}
	if len(os.Args) != 3{
		fmt.Println("参数不正确,无法校验")
		os.Exit(0)
	}

	// 检查参数是否是目录
	srcf,_ := os.Stat(os.Args[1])
	destf,_ := os.Stat(os.Args[2])

	// 捕获异常的参数 进行提醒
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("参数异常, 请检查参数正确性")
		}
	}()


	if srcf.IsDir() != true && destf.IsDir() != true  {
		fmt.Println("src参数不是目录")
		fmt.Println("dest参数不是目录")
		data := fileDiffs(os.Args[1], os.Args[2])
		fmt.Println(data)
		os.Exit(0)
	}




	var filediff FileDiff
	filediff = Diff{}
	//diffret,_ := filediff.Md5Sum("E:\\golang项目\\update","E:\\golang项目\\CICD")

	diffret,_ := filediff.Md5Sum(os.Args[1],os.Args[2])
	if len(diffret) > 1 {
		fmt.Println("升级文件有差异请检查")
		fmt.Println("异常列表:",diffret)
	}else {
		fmt.Println("文件对比正常")
	}
}


