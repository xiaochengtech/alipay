package alipay

import (
	"errors"
	"net/url"
)

// 支付通知的处理
func (c *Client) NotifyPay(raw_url string) (err error) {
	// 转换url对象
	urlObj, err := url.Parse(raw_url)
	if err != nil {
		return
	}
	// 解析查询部分
	queryMap, err := url.ParseQuery(urlObj.RawQuery)
	if err != nil {
		return
	}
	// 获取回传的数字签名并排除数字签名
	sign := ""
	if value, ok := queryMap["sign"]; ok {
		sign = value[0]
		delete(queryMap, "sign")
	} else {
		err = errors.New("没有回传数字签名")
		return
	}
	// 排除空值的对象
	body := make(map[string]interface{})
	for k, v := range queryMap {
		if v == nil || (len(v) == 0 && len(v[0]) == 0) {
			continue
		}
		body[k] = v[0]
	}
	// 验签
	if err = c.verifySignAyn(body, sign); err != nil {
		return
	}
	return
}
