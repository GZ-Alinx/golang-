package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

/*
在请求第三方接口时需要使用到的restful风格的 接口

开发核心： 编写客户端,请求服务端.

核心参数
	url  请求地址
	data post请求提交的数据
	contentType 请求体格式
	content 接口返回的内容


使用上一节示例中的Bind来充当服务端，所以要运行Bind这一节的服务

*/


// 定义第三方接口请求数据结构
type UserAPI struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// 定义第三方数据返回结构
type TempData struct {
	Msg string `json:"msg"`
	Data string `json:"data"`
}

// 客户端提交的数据结构体
type ClientRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Other interface{} `json:"other"`
}

// 客户端请求后返回的数据结构
type ClientResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}



func main() {
	//testAPI()
	r := gin.Default()
	r.POST("/getOtherAPI", getOtherAPI) //设置响应的函数
	r.Run(":9091")
}


func getOtherAPI(context *gin.Context) {
	// 将客户端提交的数据绑定到接口提ClientRequest中
	var requestData ClientRequest
	var response ClientResponse
	// 绑定接受的结构体数据到对象中
	err := context.Bind(&requestData)
	if err != nil {
		response.Code=http.StatusBadRequest
		response.Msg="请求参数错误"
		response.Data=err
		context.JSON(http.StatusBadRequest, response)
		return
	}

	// 请求第三方API接口数据
	url := "http://127.0.0.1:9090/login"
	user := UserAPI{requestData.UserName,requestData.Password}
	data,err := getRestfulAPI(url, user,"application/json")


	// 响应结构体 绑定请求数据
	var temp TempData
	json.Unmarshal(data, &temp)
	fmt.Println(temp.Msg, temp.Data)
	response.Code=http.StatusOK
	response.Msg="请求数据成功"
	response.Data=temp
	context.JSON(http.StatusOK, response)

}


func getRestfulAPI(url string, data interface{}, contentType string) ([]byte, error){
	// 创建调用API的接口Client
	client := &http.Client{Timeout: 5*time.Second}

	// 将数据转换为json
	jsonStr,_ := json.Marshal(data)

	// 将切片转换为字符串类型 bytes.NewBuffer(jsonStr)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("调用API接口出现了错误!")
		return nil,err
	}

	//resp shi Post请求的返回提,通过返回提可以提取响应数据
	res, err := ioutil.ReadAll(resp.Body)
	return res,err
}

// 测试调用方法
func testAPI(){
	url := "http://127.0.0.1:9090/login"
	user := UserAPI{"alinx", "1234"} // 请求参数

	// 调用请求接口 将参数列入方法中
	data, err := getRestfulAPI(url, user, "application/json")
	fmt.Println(data,err)

	// 拿到返回提的结构体将数据进行反序列化
	var temp TempData
	json.Unmarshal(data, &temp) // 将data数据转换为temp结构
	fmt.Println(temp.Msg, temp.Data)
}
