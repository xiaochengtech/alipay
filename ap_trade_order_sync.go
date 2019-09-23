package alipay

import (
	"encoding/json"
)

// 支付宝订单信息同步接口
func (c *Client) OrderSyncTrade(body OrderSyncTradeBody) (aliRsp OrderSyncTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.orderinfo.sync", body, params, false)
	if err != nil {
		return
	}
	var response OrderSyncTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type OrderSyncTradeBody struct {
	TradeNo       string `json:"trade_no"`                  // 支付宝订单号
	OrigRequestNo string `json:"orig_request_no,omitempty"` // 原始业务请求单号
	OutRequestNo  string `json:"out_request_no"`            // 标识一笔交易多次请求，同一笔交易多次信息同步时需要保证唯一
	BizType       string `json:"biz_type"`                  // 交易信息同步对应的业务类型
	OrderBizInfo  string `json:"order_biz_info,omitempty"`  // 商户传入同步信息
}

type OrderSyncTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo     string `json:"trade_no"`               // 支付宝交易号
	OutTradeNo  string `json:"out_trade_no,omitempty"` // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	BuyerUserId string `json:"buyer_user_id"`          // 买家在支付宝的用户id
}

type OrderSyncTradeResponseModel struct {
	Data OrderSyncTradeResponse `json:"alipay_trade_orderinfo_sync_response"` // 返回值信息
	Sign string                 `json:"sign"`                                 // 签名，参见https://docs.open.alipay.com/291/106074
}
