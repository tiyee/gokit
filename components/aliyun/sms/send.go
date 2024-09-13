package sms

import (
	"encoding/json"
	"errors"
	"github.com/tiyee/gokit/components/aliyun/internal/openapi"
)

func (c *Client) Send(phoneNumbers, signName, templateCode string, templateParam string) (*Response, error) {
	httpMethod := "POST"
	xAcsAction := "SendSms"
	openApi := openapi.NewOpenApi(c)
	query := map[string]string{
		"PhoneNumbers":  phoneNumbers,
		"SignName":      signName,
		"TemplateCode":  templateCode,
		"TemplateParam": templateParam,
	}
	body, err := openApi.CallApi(httpMethod, c.CanonicalUri(), c.Host(), xAcsAction, c.xAcsVersion, query)
	if err != nil {
		return nil, err
	}
	var resp Response
	defer body.Close()
	if err = json.NewDecoder(body).Decode(&resp); err == nil {
		if resp.Code != "OK" {
			err = errors.New(resp.Message)
		}
		return &resp, nil
	}
	return nil, err
}
