package direct_mail

import (
	"github.com/tiyee/gokit/slice"
	"net/smtp"
	"strings"
	"time"
)

type Client struct {
	authUser       string
	authPasswd     string
	host           string
	addr           string
	replyToAddress string
}
type Option func(*Client)

func New(user, passwd, addr string, opts ...Option) *Client {
	c := &Client{
		authUser:       user,
		authPasswd:     passwd,
		addr:           addr,
		host:           strings.Split(addr, ":")[0],
		replyToAddress: user, // default auth user
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
func (c *Client) smtpAuth() smtp.Auth {
	return smtp.PlainAuth("", c.authUser, c.authPasswd, c.host)
}

// SendEmail 发送邮件
// to: 收件人地址
// cc: 抄送地址
// bcc: 密送地址
func (c *Client) SendEmail(subject, mailType, body string, to, cc, bcc []string) error {
	auth := c.smtpAuth()
	ccAddress := strings.Join(cc, ";")
	bccAddress := strings.Join(bcc, ";")
	toAddress := strings.Join(to, ";")
	date := time.Now().Format(time.RFC1123Z)
	contentType := "Content-Type: text/plain; charset=UTF-8"
	if strings.ToLower(mailType) == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	}
	replyToAddress := c.replyToAddress
	msg := []byte("To: " + toAddress + "\r\nFrom: " + c.authUser + "\r\nSubject: " + subject + "\r\nDate: " + date + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + ccAddress + "\r\nBcc: " + bccAddress + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := slice.Merge(to, cc, bcc)
	err := smtp.SendMail(c.addr, auth, c.authUser, sendTo, msg)
	return err
}
