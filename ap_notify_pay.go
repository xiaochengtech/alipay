package alipay

import (
	"encoding/json"
	"errors"
	"net/url"
)

// 支付通知的处理
func (c Client) NotifyPay(bodyStr string) (params NotifyPayParams, origin map[string]interface{}, err error) {
	// 解析并验证签名
	origin, err = c.payNotifyParseParams(bodyStr, &params)
	if err != nil {
		return
	}
	// 验证notify_id
	if len(params.NotifyId) == 0 {
		err = errors.New("不存在notify_id")
		return
	}
	return
}

// 支付结果通知的参数
type NotifyPayParams struct {
	// 通知参数
	NotifyType string `json:"notify_type"` // 通知类型，参见constant.go
	NotifyId   string `json:"notify_id"`   // 91722adff935e8cfa58b3aabf4dead6ibe
	NotifyTime string `json:"notify_time"` // 2017-02-16 21:46:15
	// 公共参数
	AppId     string `json:"app_id"`      // 应用ID
	AuthAppId string `json:"auth_app_id"` // 授权商户应用ID
	Chatset   string `json:"charset"`     // 字符集，参见constant.go
	Version   string `json:"version"`     // 接口版本，参见constant.go
	// 业务参数
	TradeNo      string `json:"trade_no"`       // 支付宝交易号
	OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
	BuyerLogonId string `json:"buyer_logon_id"` // 买家支付宝账号
	TradeStatus  string `json:"trade_status"`   // 交易状态
	TotalAmount  string `json:"total_amount"`   // 订单总金额
}

// 验证签名
func (c Client) payNotifyParseParams(bodyStr string, params *NotifyPayParams) (origin map[string]interface{}, err error) {
	// 解析查询部分
	queryMap, err := url.ParseQuery(bodyStr)
	if err != nil {
		return
	}
	// 获取回传的数字签名并排除数字签名
	sign := ""
	if value, ok := queryMap["sign"]; ok {
		sign = value[0]
		delete(queryMap, "sign")
		delete(queryMap, "sign_type")
	} else {
		err = errors.New("没有回传数字签名")
		return
	}
	// 排除空值的对象
	origin = make(map[string]interface{})
	for k, v := range queryMap {
		if v == nil || (len(v) == 0 && len(v[0]) == 0) {
			continue
		}
		origin[k] = v[0]
	}
	// 验签
	if err = c.verifySignAyn(origin, sign); err != nil {
		return
	}
	// 解析Json对象
	jsonStr, _ := json.Marshal(origin)
	err = json.Unmarshal(jsonStr, params)
	return
}
