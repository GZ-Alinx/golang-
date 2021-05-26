package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
	"log"
	"strconv"
	"sync"
)

func Mail(mailTo string, subject string, body string, wg *sync.WaitGroup) error {
	defer wg.Done()
	// 服务器smtp信息，登录用户和密码
	mailConn := map[string]string{
		"username": "itadminlx",
		"authCode": "****",
		"host":     "smtp.163.com",
		"port":     "465",
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "itadminlx@163.com")
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// 拼接邮件信息
	port, _ := strconv.Atoi(mailConn["port"])
	d := gomail.NewDialer(mailConn["host"], port, mailConn["username"], mailConn["authCode"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		log.Fatalln("To", mailTo, "Send Email Faild", err)
		return fmt.Errorf("faild")
	} else {
		log.Fatalln("To", mailTo, "Send Email Successfully!")
		return fmt.Errorf("success")
	}
}

func Send() {
	var wg sync.WaitGroup
	mailTo := []string{
		"itadminlx@163.com",
	}
	subject := "主题"
	body := "主体信息"
	for _, mail := range mailTo {
		wg.Add(1)
		go Mail(mail, subject, body, &wg)
	}
}

func main() {
	Send()
}
