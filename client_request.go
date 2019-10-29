package alipay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
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
func (c *Client) doAlipay(method string, body interface{}, params BodyMap, isGBK bool, isPost bool) (bytes []byte, err error) {
	// 获取请求参数
	urlParam, err := c.doParams(method, body, params, isPost)
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
	if isPost {
		resp, err = client.Post(reqUrl, "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(urlParam))
	} else {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, urlParam)
		resp, err = client.Get(reqUrl)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if isGBK {
		bytes, err = simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	}
	return
}

// 生成请求参数
func (c *Client) doParams(method string, body interface{}, params BodyMap, isPost bool) (urlParam string, err error) {
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
		params["charset"] = "utf-8"
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
	// 将Body参数转换为JSON字符串
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	if isPost {
		params["biz_content"] = string(bodyStr)
	} else {
		var bodyParams BodyMap
		if err = json.Unmarshal(bodyStr, &bodyParams); err != nil {
			return
		}
		for k, v := range bodyParams {
			params[k] = v
		}
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

// 格式化请求URL参数
func (c *Client) FormatURLParam(body BodyMap) string {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
