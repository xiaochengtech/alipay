package alipay

import (
	"encoding/json"
)

// 统一收单交易结算接口
func (c Client) TradeOrderSettle(body TradeOrderSettleBody) (aliRsp TradeOrderSettleResponse, err error) {
	params := BodyMap{
		"biz_content": c.GenerateBizContent(body),
	}
	bytes, err := c.doAlipay("alipay.trade.order.settle", params)
	if err != nil {
		return
	}
	var response TradeOrderSettleResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type TradeOrderSettleBody struct {
	OutRequestNo      string             `json:"out_request_no"`        // 结算请求流水号
	TradeNo           string             `json:"trade_no"`              // 支付宝订单号
	RoyaltyParameters []RoyaltyParameter `json:"royalty_parameters"`    // 分账明细信息
	OperatorId        string             `json:"operator_id,omitempty"` // 商户操作员编号
}

type TradeOrderSettleResponse struct {
	ResponseModel
	// 响应参数
	TradeNo string `json:"trade_no"` // 支付宝交易号
}

type TradeOrderSettleResponseModel struct {
	Data TradeOrderSettleResponse `json:"alipay_trade_order_settle_response"` // 返回值信息
	Sign string                   `json:"sign"`                               // 签名，参见https://docs.open.alipay.com/291/106074
}
