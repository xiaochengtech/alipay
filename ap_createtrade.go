package alipay

import (
	"encoding/json"
)

// 统一收单交易创建接口
func (c *Client) CreateTrade(body CreateTradeBody, notifyUrl string) (aliRsp CreateTradeResponse, err error) {
	params := BodyMap{}
	if notifyUrl != "" {
		params["notify_url"] = notifyUrl
	}
	bytes, err := c.doAliPay("alipay.trade.create", body, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &aliRsp)
	return
}

type CreateTradeBody struct {
	OutTradeNo string `json:"out_trade_no"`        // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	SellerId   string `json:"seller_id,omitempty"` // 卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	// TODO 其他参数https://docs.open.alipay.com/api_1/alipay.trade.create/
}

type CreateTradeResponse struct {
	ResponseModel
	// 响应参数
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
}
