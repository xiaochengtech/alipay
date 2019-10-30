package alipay

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
}

// 向支付宝发送Post请求
func (c *Client) doAlipay(method string, params BodyMap) (bytes []byte, err error) {
	// 获取请求参数
	urlParam, err := c.doParams(method, params)
	if err != nil {
		return
	}
	// 发起请求
	var (
		reqUrl string
		resp   *http.Response
	)
	if c.isProd {
		reqUrl = baseUrl
	} else {
		reqUrl = baseUrlSandbox
	}
	resp, err = client.Post(reqUrl, "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(urlParam))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	return
}

// 生成请求参数
func (c *Client) doParams(method string, params BodyMap) (urlParam string, err error) {
	// 生成公共请求参数
	// notify_url按需提前传入至params
	params["app_id"] = c.config.AppId
	params["method"] = method
	if c.config.Format != "" {
		params["format"] = c.config.Format
	} else {
		params["format"] = FormatJson
	}
	if c.config.Charset != "" {
		params["charset"] = c.config.Charset
	} else {
		params["charset"] = CharSetUTF8
	}
	var signType string
	if c.config.SignType != "" {
		signType = c.config.SignType
	} else {
		signType = SignTypeRSA2
	}
	params["sign_type"] = signType
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	if c.config.Version != "" {
		params["version"] = c.config.Version
	} else {
		params["version"] = Version1
	}
	if c.config.AppAuthToken != "" {
		params["app_auth_token"] = c.config.AppAuthToken
	}
	// 获取签名
	pKey := c.FormatPrivateKey(c.privateKey)
	sign, err := c.getSign(params, signType, pKey)
	if err != nil {
		return
	}
	params["sign"] = sign
	// 格式化请求URL参数
	urlParam = c.FormatURLParam(params)
	return
}

// 生成业务字段
func (c *Client) GenerateBizContent(body interface{}) string {
	bodyStr, _ := json.Marshal(body)
	return string(bodyStr)
}

// 格式化请求URL参数
func (c *Client) FormatURLParam(body BodyMap) string {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}

// 生成到BodyMap中
func (c *Client) ConvertToBodyMap(params interface{}) (body BodyMap) {
	paramStr, _ := json.Marshal(params)
	_ = json.Unmarshal(paramStr, &body)
	return
}
