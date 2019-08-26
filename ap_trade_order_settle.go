package alipay

import (
	"encoding/json"
)

// 统一收单交易结算接口
func (c *Client) OrderSettleTrade(body OrderSettleTradeBody) (aliRsp OrderSettleTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.order.settle", body, params, false)
	if err != nil {
		return
	}
	var response OrderSettleTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	/* if err = c.verifySignSync(response.Data, response.Sign); err != nil {
		return
	} */
	aliRsp = response.Data
	return
}

type OrderSettleTradeBody struct {
	OutRequestNo      string             `json:"out_request_no"`        // 结算请求流水号
	TradeNo           string             `json:"trade_no"`              // 支付宝订单号
	RoyaltyParameters []RoyaltyParameter `json:"royalty_parameters"`    // 分账明细信息
	OperatorId        string             `json:"operator_id,omitempty"` // 商户操作员编号
}

type OrderSettleTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo string `json:"trade_no"` // 支付宝交易号
}

type OrderSettleTradeResponseModel struct {
	Data OrderSettleTradeResponse `json:"alipay_trade_order_settle_response"` // 返回值信息
	Sign string                   `json:"sign"`                               // 签名，参见https://docs.open.alipay.com/291/106074
}
