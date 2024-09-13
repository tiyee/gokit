package aliyun

import (
	"fmt"
	"github.com/tiyee/gokit/components/aliyun/sms"
)

func SendSms() {
	client := sms.New("xxxxxxxx", "xxxxxxxxx")
	resp, err := client.Send("18600000000", "XXXX", "SMS_0000000000", "{\"code\": \"1234\"}")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
