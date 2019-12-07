package mail

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo string, title string, content string) error {
	return SendMails([]string{mailTo}, title, content)
}

//SendMail
func SendMails(mailTo []string, title string, content string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "971907847@qq.com",
		"pass": "izktvwlunmzbbfbj",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	//发件人
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", title)                           //设置邮件主题
	m.SetBody("text/html", content)                         //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	return d.DialAndSend(m)
}
