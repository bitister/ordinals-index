package utils

import (
	"net/smtp"

	"github.com/astaxie/beego"
	"github.com/jordan-wright/email"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = ":587"
	SMTPUsername = "thespacerace.tsr@gmail.com"
	SMTPPassword = "xqufzreqindiyrex"
	MaxClient    = 50
)

// func SendCodeEmail(receiver, code string) {
// 	pool, err := email.NewPool(SMTPHost+SMTPPort, MaxClient, smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost))
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	e := &email.Email{
// 		From:    SMTPUsername,
// 		To:      []string{receiver},
// 		Subject: "验证码",
// 		Text:    []byte("【TSR】" + code),
// 	}
// 	err = pool.Send(e, 5*time.Second)
// 	defer pool.Close()
// 	if err != nil {
// 		beego.Error(err)
// 	}
// }

func SendCodeEmail(receiver, code string) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	e := &email.Email{
		From:    SMTPUsername,
		To:      []string{receiver},
		Subject: "验证码",
		Text:    []byte("【TSR】" + code),
	}
	err := e.Send(SMTPHost+SMTPPort, auth)
	if err != nil {
		beego.Error(err)
	}
}

func sendHtmlEmail(receiver string, html []byte) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	e := &email.Email{
		From:    SMTPUsername,
		To:      []string{receiver},
		Subject: "Verification Code",
		HTML:    html,
	}
	err := e.Send(SMTPHost+SMTPPort, auth)
	if err != nil {
		beego.Error(err)
	}
}
