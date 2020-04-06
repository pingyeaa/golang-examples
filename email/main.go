package main

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	//e := &email.Email{
	//	From:    "平也 <pingyeaa@163.com>",
	//	To:      []string{"xxxxxxx@qq.com"},
	//	Subject: "发现惊天大秘密！",
	//	Text:    []byte("平也好帅好有智慧哦~"),
	//}

	//e.From = "pingyeaa@163.com"
	e.From = "平也 <pingyeaa@163.com>"
	e.To = []string{"xxxxx@qq.com"}
	e.Subject = "发现惊天大秘密！"
	e.Text = []byte("平也好帅好有智慧哦~")
	//e.Cc = []string{"xxxxxxx@qq.com"}
	//e.Bcc = []string{"xxxxxxx@qq.com"}
	e.AttachFile("attachment.txt")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "pingyeaa@163.com", "<你的密码>", "smtp.163.com"))
	if err != nil {
		panic(err)
	}
}
