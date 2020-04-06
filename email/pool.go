package main

import (
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	auth := smtp.PlainAuth("", "pingyeaa@163.com", "<你的密码>", "smtp.163.com")
	p, _ := email.NewPool("smtp.163.com:25", 4, auth)

	e := email.NewEmail()
	e.From = "平也 <pingyeaa@163.com>"
	e.To = []string{"xxxx@qq.com"}
	e.Subject = "发现惊天大秘密！"
	e.Text = []byte("平也好帅好有智慧哦~")

	p.Send(e, 10*time.Second)
}
