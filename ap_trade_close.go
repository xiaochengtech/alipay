package alipay

import (
	"encoding/json"
)

// 统一收单交易关闭接口
func (c *Client) CloseTrade(body CloseTradeBody) (aliRsp CloseTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.close", body, params, false)
	if err != nil {
		return
	}
	var response CloseTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type CloseTradeBody struct {
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
	OperatorId string `json:"operator_id,omitempty"`  // 卖家端自定义的的操作员 ID
}

type CloseTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
}

type CloseTradeResponseModel struct {
	Data CloseTradeResponse `json:"alipay_trade_close_response"` // 返回值信息
	Sign string             `json:"sign"`                        // 签名，参见https://docs.open.alipay.com/291/106074
}
