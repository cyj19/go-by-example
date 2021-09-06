package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
)

/*
	邮件功能开发步骤:
	1. 确定使用邮件协议
	2. 编写配置
	3. 创建收件人、主题、正文
	4. 发送邮件

	smtp包发送邮件注意点：
	1. 要想主题、发送人、收件人在邮件上显示需要自己构建符合规则的字符串
	2. \r\n为固定换行格式
	3. 多个收件人需要在构建收件人字符串时以分号分隔，且在SendMail的address中也要添加相应的收件人
	4. 附件需要指定Content-Type参数,并指定一个分隔符,将邮件头,正文部分与附件分隔开,分隔符可以任意设置但不要是中文
	   content_type := "Content-Type: multipart/mixed; boundary=分隔符\r\n"
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
	// 构建邮件头部信息
	from := fmt.Sprintf("From:%s\r\n", SMTP_MAIL_USER)
	var to string
	for _, val := range address {
		to += val + ";"
	}
	to = fmt.Sprintf("To:%s\r\n", to)
	sub := fmt.Sprintf("Subject:%s\r\n", subject)
	// 分隔符
	delimiter := "module"
	content_type := "Content-Type: multipart/mixed; boundary=" + delimiter + "\r\n"
	header := from + to + sub + content_type
	// 通过分隔符分隔头信息和正文
	mailContent := header + fmt.Sprintf("\r\n--%s\r\n", delimiter)
	mailContent += content
	// 分隔正文和附件信息
	fileheader := fmt.Sprintf("\r\n--%s\r\n", delimiter)
	path := "./a.txt"
	name := "a.txt"
	fileheader += "Content-Type: application/octet-stream\r\n"
	fileheader += "Content-Transfer-Encoding: base64\r\n"
	fileheader += "Content-Disposition: attachment; filename=\"" + name + "\"\r\n\r\n"
	// 读取附件内容并编码
	fileinfo, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	filecontent := base64.StdEncoding.EncodeToString(fileinfo)
	filemessage := fileheader + filecontent
	// 分隔符结束
	mailContent += filemessage + "\r\n--" + delimiter + "--\r\n\r\n"
	fmt.Println(mailContent)
	msg := []byte(mailContent)
	// 发送邮件
	return smtp.SendMail(addr, auth, SMTP_MAIL_USER, address, msg)
}
