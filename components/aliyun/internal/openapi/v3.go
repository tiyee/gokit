package openapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/tiyee/gokit/slice"
	"github.com/tiyee/gokit/strlib"
	"io"
	"sort"

	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// @link https://help.aliyun.com/zh/sdk/product-overview/v3-request-structure-and-signature
type Request struct {
	httpMethod   string
	scheme       string
	url          string
	canonicalUri string
	host         string
	xAcsAction   string
	xAcsVersion  string
	headers      map[string]string
	body         string
	queryParam   map[string]string
}

func NewRequest(httpMethod, scheme, canonicalUri, host, xAcsAction, xAcsVersion string) *Request {
	req := &Request{
		httpMethod:   httpMethod,
		scheme:       scheme,
		canonicalUri: canonicalUri,
		host:         host,
		xAcsAction:   xAcsAction,
		xAcsVersion:  xAcsVersion,
		headers:      make(map[string]string),
		queryParam:   make(map[string]string),
	}
	req.headers["host"] = host
	req.headers["x-acs-action"] = xAcsAction
	req.headers["x-acs-version"] = xAcsVersion
	req.headers["x-acs-date"] = time.Now().UTC().Format(time.RFC3339)
	req.headers["x-acs-signature-nonce"] = strlib.NonceStr()
	return req
}

type IBaseConfig interface {
	AccessKeyId() string
	AccessKeySecret() string
	Scheme() string
}
type OpenApiV3 struct {
	AccessKeyId     string
	AccessKeySecret string
	ALGORITHM       string
	Scheme          string
}

func NewOpenApi(getter IBaseConfig) OpenApiV3 {
	return OpenApiV3{
		ALGORITHM:       "ACS3-HMAC-SHA256",
		AccessKeyId:     getter.AccessKeyId(),
		Scheme:          getter.Scheme(),
		AccessKeySecret: getter.AccessKeySecret()}

}

func (c OpenApiV3) CallApi(httpMethod, canonicalUri, host, xAcsAction, xAcsVersion string, queryParam map[string]string) (io.ReadCloser, error) {
	req := NewRequest(httpMethod, c.Scheme, canonicalUri, host, xAcsAction, xAcsVersion)
	for k, v := range queryParam {
		req.queryParam[k] = v
	}
	c.getAuthorization(req)

	// 调用API
	return callAPI(req)

}

func callAPI(req *Request) (io.ReadCloser, error) {
	urlStr := req.scheme + "://" + req.host + req.canonicalUri
	q := url.Values{}
	keys := slice.Keys(req.queryParam)
	sort.Strings(keys)
	for _, k := range keys {
		v := req.queryParam[k]
		q.Set(k, v)
	}
	urlStr += "?" + q.Encode()

	httpReq, err := http.NewRequest(req.httpMethod, urlStr, strings.NewReader(req.body))
	if err != nil {
		return nil, err
	}

	for key, value := range req.headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c OpenApiV3) getAuthorization(req *Request) {
	canonicalQueryString := ""
	keys := slice.Keys(req.queryParam)
	sort.Strings(keys)
	for _, k := range keys {
		v := req.queryParam[k]
		canonicalQueryString += percentCode(url.QueryEscape(k)) + "=" + percentCode(url.QueryEscape(v)) + "&"
	}
	canonicalQueryString = strings.TrimSuffix(canonicalQueryString, "&")
	//fmt.Printf("canonicalQueryString========>%s\n", canonicalQueryString)

	hashedRequestPayload := sha256Hex(req.body)
	req.headers["x-acs-content-sha256"] = hashedRequestPayload

	canonicalHeaders := ""
	signedHeaders := ""
	HeadersKeys := slice.Keys(req.headers)
	sort.Strings(HeadersKeys)
	for _, k := range HeadersKeys {
		lowerKey := strings.ToLower(k)
		if lowerKey == "host" || strings.HasPrefix(lowerKey, "x-acs-") || lowerKey == "content-type" {
			canonicalHeaders += lowerKey + ":" + req.headers[k] + "\n"
			signedHeaders += lowerKey + ";"
		}
	}
	signedHeaders = strings.TrimSuffix(signedHeaders, ";")

	canonicalRequest := req.httpMethod + "\n" + req.canonicalUri + "\n" + canonicalQueryString + "\n" + canonicalHeaders + "\n" + signedHeaders + "\n" + hashedRequestPayload
	//fmt.Printf("canonicalRequest========>\n%s\n", canonicalRequest)

	hashedCanonicalRequest := sha256Hex(canonicalRequest)
	stringToSign := c.ALGORITHM + "\n" + hashedCanonicalRequest
	//fmt.Printf("stringToSign========>\n%s\n", stringToSign)

	byteData, err := hmac256([]byte(c.AccessKeySecret), stringToSign)
	if err != nil {
		fmt.Println(err)
	}
	signature := strings.ToLower(hex.EncodeToString(byteData))

	authorization := c.ALGORITHM + " Credential=" + c.AccessKeyId + ",SignedHeaders=" + signedHeaders + ",Signature=" + signature
	//fmt.Printf("authorization========>%s\n", authorization)
	req.headers["Authorization"] = authorization
}

func hmac256(key []byte, toSignString string) ([]byte, error) {
	// 实例化HMAC-SHA256哈希
	h := hmac.New(sha256.New, key)
	// 写入待签名的字符串
	_, err := h.Write([]byte(toSignString))
	if err != nil {
		return nil, err
	}
	// 计算签名并返回
	return h.Sum(nil), nil
}

func sha256Hex(str string) string {
	// 实例化SHA-256哈希函数
	hash := sha256.New()
	// 将字符串写入哈希函数
	_, _ = hash.Write([]byte(str))
	// 计算SHA-256哈希值并转换为小写的十六进制字符串
	hexString := hex.EncodeToString(hash.Sum(nil))

	return hexString
}

func percentCode(str string) string {
	// 替换特定的编码字符
	str = strings.ReplaceAll(str, "+", "%20")
	str = strings.ReplaceAll(str, "*", "%2A")
	str = strings.ReplaceAll(str, "%7E", "~")
	return str
}
func main() {
	// RPC接口请求
	httpMethod := "POST"
	canonicalUri := "/"
	host := "ecs.cn-beijing.aliyuncs.com"
	xAcsAction := "DescribeInstances"
	xAcsVersion := "2014-05-26"
	req := NewRequest(httpMethod, "", canonicalUri, host, xAcsAction, xAcsVersion)
	req.queryParam["RegionId"] = "cn-beijing"
	req.queryParam["VpcId"] = "vpc-2zeo42r27y4opXXXXXXXX"

	/*    // ROA接口POST请求
	httpMethod := "POST"
	canonicalUri := "/clusters"
	host := "cs.cn-beijing.aliyuncs.com"
	xAcsAction := "CreateCluster"
	xAcsVersion := "2015-12-15"
	req := NewRequest(httpMethod, canonicalUri, host, xAcsAction, xAcsVersion)
	body := make(map[string]string)
	body["name"] = "testDemo"
	body["region_id"] = "cn-beijing"
	body["cluster_type"] = "Kubernetes"
	body["vpcid"] = "vpc-2zeo42r27y4opXXXXXXXX"
	body["service_cidr"] = "172.16.1.0/20"
	body["security_group_id"] = "sg-2zec0dm6qi66XXXXXXXX"
	jsonBytes, err := json.Marshal(body)
	if err != nil {
	fmt.Println("Error marshaling to JSON:", err)
	return
	}
	req.body = string(jsonBytes)
	req.headers["content-type"] = "application/json; charset=utf-8"*/

	/*    // ROA接口GET请求
	httpMethod := "GET"
	            // canonicalUri如果存在path参数，需要对path参数encode，percentCode({path参数})
	canonicalUri := "/clusters/" + percentCode("cb7cd6b9bde934f6193801878XXXXXXXX") + "/resources"
	host := "cs.cn-beijing.aliyuncs.com"
	xAcsAction := "DescribeClusterResources"
	xAcsVersion := "2015-12-15"
	req := NewRequest(httpMethod, canonicalUri, host, xAcsAction, xAcsVersion)
	req.queryParam["with_addon_resources"] = "true"*/

	// 签名过程

}
