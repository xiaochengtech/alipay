package alipay

import (
	"encoding/json"
)

// 统一收单交易撤销接口
func (c Client) TradeCancel(body TradeCancelBody) (aliRsp TradeCancelResponse, err error) {
	params := BodyMap{
		"biz_content": c.GenerateBizContent(body),
	}
	bytes, err := c.doAlipay("alipay.trade.cancel", params)
	if err != nil {
		return
	}
	var response TradeCancelResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type TradeCancelBody struct {
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
}

type TradeCancelResponse struct {
	ResponseModel
	// 响应参数
	TradeNo            string `json:"trade_no"`                       // 支付宝交易号
	OutTradeNo         string `json:"out_trade_no"`                   // 商户订单号
	RetryFlag          string `json:"retry_flag"`                     // 是否需要重试
	Action             string `json:"action"`                         // 本次撤销触发的交易动作
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`       // 返回的退款时间
	RefundSettlementId string `json:"refund_settlement_id,omitempty"` // 返回的退款清算编号
}

type TradeCancelResponseModel struct {
	Data TradeCancelResponse `json:"alipay_trade_cancel_response"` // 返回值信息
	Sign string              `json:"sign"`                         // 签名，参见https://docs.open.alipay.com/291/106074
}
