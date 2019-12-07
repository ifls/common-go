package mail

import (
	"log"
	"testing"
)

func TestSend(t *testing.T) {
	//定义收件人
	mailTo := "coolday3741@163.com"
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "Good 2"
	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
	}
}
