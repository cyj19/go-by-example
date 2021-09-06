package main

import (
	"fmt"
	"net/smtp"
)

/*
	邮件功能开发步骤:
	1. 确定使用邮件协议
	2. 编写配置
	3. 创建收件人、主题、正文
	4. 发送邮件
*/

const (
	SMTP_MAIL_HOST = "smtp.163.com"   // 邮件服务地址
	SMTP_MAIL_PORT = "465"            // 端口
	SMTP_MAIL_USER = "sender@163.com" // 发件人
	SMTP_MAIL_PWD  = "xxxxxxx"        // 授权密码
)

func main() {
	//收件人
	address := []string{"address1@163.com"}
	// 主题
	subject := "test mail"
	// 正文
	content := "this is the email body"
	// 发送邮件
	sendMail(address, subject, content)
}

func sendMail(address []string, subject string, content string) error {
	// 认证
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	// smtp服务器地址
	addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
	// 消息
	s := fmt.Sprintf("To:%s\r\nFrom:%s\r\nSubject:%s\r\n\r\n%s", address[0], SMTP_MAIL_USER, subject, content)
	msg := []byte(s)
	// 发送邮件
	return smtp.SendMail(addr, auth, SMTP_MAIL_USER, address, msg)
}
