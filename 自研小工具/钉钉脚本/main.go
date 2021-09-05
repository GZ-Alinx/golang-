package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Dingding interface {
	Texts(token, texts, users string) error
	Sends(token, texts, users string) error
}

type Alert struct{}

func (info *Alert) Texts(token, texts, users string) error {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
	method := "POST"

	// 注意携带机器人中设置的关键字 message
	RequestsData := fmt.Sprintf(`
	{
		"at": {
			"atMobiles": [
	        "%s"
	    ],
		"isAtAll": false},
		"text": {
		"content": "message： %s"
		},
		"msgtype": "text"
	}`, texts, users)
	fmt.Println(RequestsData)
	payload := strings.NewReader(RequestsData)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (info *Alert) Markdown(token, texts, users string) error {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
	method := "POST"
	// 此方法 缺少参数  官方文档地址： https://developers.dingtalk.com/document/robots/custom-robot-access
	payload := strings.NewReader(RequestsData)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func main() {
	var dingding Dingding
	dingding = new(Alert) // 指针引用
	// dingding = Alert{}    // 值类型引用
	dingding.Texts("xxx", "xxx", "测试@特定人员")
}
