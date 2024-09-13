package aliyun

import (
	"fmt"
	"github.com/tiyee/gokit/components/aliyun/direct_mail"
)

const (
	EmailUser     = "no-reply@demo.com"
	EmailPassword = "xxxxxxxxx"
	EmailAddr     = "smtpdm.aliyun.com:80"
	EmailSubject  = "Demo | Verify Your Email"
)

func SendEmail() {
	client := direct_mail.New(EmailUser, EmailPassword, EmailAddr)
	body := fmt.Sprintf("<html><body><p>Enter this code to sign up and join unlimited 3D creative universe:</p><strong>%s</strong></body></html>", "123")

	err := client.SendEmail(EmailSubject, "html", body, []string{"tiyee@xxx.com"}, []string{}, []string{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("success")
}
