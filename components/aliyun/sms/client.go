package sms

type Client struct {
	accessKeyId     string
	accessKeySecret string
	canonicalUri    string
	host            string
	scheme          string
	xAcsVersion     string
}
type Option func(*Client)

func New(accessKeyId, accessKeySecret string, opts ...Option) *Client {
	client := &Client{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		scheme:          "https",
		canonicalUri:    "/",
		host:            "dysmsapi.aliyuncs.com",
		xAcsVersion:     "2017-05-25",
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}
func (c *Client) AccessKeyId() string {
	return c.accessKeyId
}
func (c *Client) AccessKeySecret() string {
	return c.accessKeySecret
}
func (c *Client) XAcsVersion() string {
	return c.xAcsVersion
}
func (c *Client) Host() string {
	return c.host
}
func (c *Client) CanonicalUri() string {
	return c.canonicalUri
}
func (c *Client) Scheme() string {
	return c.scheme
}
