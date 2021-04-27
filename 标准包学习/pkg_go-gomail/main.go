package main

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"strconv"
)

func main() {
	// 服务器smtp信息，登录用户和密码
	mailConn := map[string]string{
		"username": "itadminlx",
		"authCode": "GWVYEPUOQTWWNMBYJA",
		"host":     "smtp.163.com",
		"port":     "25",
	}

	// 设置邮箱主体信息 发件人等
	//Send()
	m := gomail.NewMessage()
	m.SetHeader("From", "itadminlx@163.com")
	m.SetHeader("To", "itadminlx@163.com", "1135189009@qq.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan") //抄送
	m.SetHeader("Subject", "subject") // 邮件标题
	m.SetBody("text/html", "body")    // 邮件内容
	// m.Attach("/home/Alex/lolcat.jpg") //附件

	// 组装邮件信息 进行拨号登录
	//d := gomail.NewDialer("smtp.163.com", 25, "itadminlx", "WVYEPUOQTWWNMBYJ")
	port, _ := strconv.Atoi(mailConn["port"])
	d := gomail.NewDialer(mailConn["host"], port, mailConn["username"], mailConn["authCode"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

//func main() {
//	var wg sync.WaitGroup
//	//mailTo := []string{
//	//	"itadminlx@163.com",
//	//}
//	//subject := "主题"
//	//body := "主体信息"
//	////for _,mail := range mailTo{
//	////	wg.Add(1)
//	////
//	////}
//	Sendmail(subject,body,&wg)
//
//}
