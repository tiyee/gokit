package offiaccount

import (
	"encoding/xml"
	"log"
	"time"
)

type BaseMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
}
type TextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	Content      string
}
type PicMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	PicUrl       string
	MediaId      string
}
type LinkMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	Title        string
	Description  string
	Url          string
}
type RespBaseMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	XMLName      xml.Name `xml:"xml"`
}
type RespTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	XMLName      xml.Name `xml:"xml"`
}

func BuildReply(from string, to string, msg string) string {
	data := RespTextMsg{
		ToUserName:   to,
		FromUserName: from,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      msg,
	}
	ret, err := xml.Marshal(&data)
	if err != nil {
		log.Println(err.Error())
		return ""

	}
	return string(ret)
}
