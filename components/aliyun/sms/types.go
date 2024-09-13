package sms

type Config struct {
	accessKeyId     string
	accessKeySecret string
}

func (c Config) AccessKeyId() string {
	return c.accessKeyId
}

func (c Config) AccessKeySecret() string {
	return c.accessKeySecret
}

type Response struct {
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	Code      string `json:"Code"`
	BizId     string `json:"BizId"`
}
