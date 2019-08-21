package alipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 生成请求参数
func (c *Client) doGenerateParams(method string, body interface{}, params BodyMap) (urlParam string, err error) {
	// 将Body参数转换为JSON字符串
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
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
	params["biz_content"] = string(bodyStr)
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

// 向支付宝发送请求
func (c *Client) doAliPay(method string, body interface{}, params BodyMap, isGBK bool) (bytes []byte, err error) {
	// 获取请求参数
	urlParam, err := c.doGenerateParams(method, body, params)
	if err != nil {
		return
	}
	// 发起请求
	var url string
	if c.isProd {
		url = baseUrl
	} else {
		url = baseUrlSandbox
	}
	resp, err := http.Post(url, "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(urlParam))
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

// 格式化请求URL参数
func (c *Client) FormatURLParam(body BodyMap) string {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
