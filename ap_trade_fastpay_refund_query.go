package alipay

import (
	"encoding/json"
	"time"
)

// 统一收单交易退款查询接口
func (c *Client) RefundQueryTrade(body RefundQueryTradeBody) (aliRsp RefundQueryTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.fastpay.refund.query", body, params, false)
	if err != nil {
		return
	}
	var response RefundQueryTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	/* if err = c.verifySignSync(response.Data, response.Sign); err != nil {
		return
	} */
	aliRsp = response.Data
	return
}

type RefundQueryTradeBody struct {
	TradeNo      string `json:"trade_no,omitempty"`     // 支付宝订单号
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	OutRequestNo string `json:"out_request_no"`         // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传
	OrgPId       string `json:"org_pid,omitempty"`      // 退款的交易所属收单机构的pid
}

type RefundQueryTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo               string              `json:"trade_no,omitempty"`                        // 支付宝交易号
	OutTradeNo            string              `json:"out_trade_no,omitempty"`                    // 商户订单号
	OutRequestNo          string              `json:"out_request_no,omitempty"`                  // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传
	RefundReason          string              `json:"refund_reason,omitempty"`                   // 退款的原因说明
	TotalAmount           string              `json:"total_amount,omitempty"`                    // 订单总金额
	RefundAmount          string              `json:"refund_amount,omitempty"`                   // 退款总金额
	RefundRoyaltys        []RefundRoyaltyInfo `json:"refund_royaltys,omitempty"`                 // 退分账明细信息
	GmtRefundPay          time.Time           `json:"gmt_refund_pay,omitempty"`                  // 退款支付时间
	FundBillList          []FundBillListInfo  `json:"refund_detail_item_list,omitempty"`         // 退款使用的资金渠道
	SendbackFee           string              `json:"send_back_fee,omitempty"`                   // 本次商户实际退回金额
	RefundSettlementId    string              `json:"refund_settlement_id,omitempty"`            // 支付清算编号
	RefundBuyerAmount     string              `json:"present_refund_buyer_amount,omitempty"`     // 本次退款金额中买家退款金额
	RefundDiscountAmount  string              `json:"present_refund_discount_amount,omitempty"`  // 本次退款金额中平台优惠退款金额
	RefundMdiscountAmount string              `json:"present_refund_mdiscount_amount,omitempty"` // 本次退款金额中商家优惠退款金额
}

type RefundRoyaltyInfo struct {
	RefundAmount string `json:"refund_amount"`             // 退分账金额
	RoyaltyType  string `json:"royalty_type,omitempty"`    // 分账类型
	ResultCode   string `json:"result_code"`               // 退分账结果码
	TransOut     string `json:"trans_out,omitempty"`       // 转出人支付宝账号对应用户ID
	TransOutType string `json:"trans_out_email,omitempty"` // 转出人支付宝账号
	TransIn      string `json:"trans_in,omitempty"`        // 转入人支付宝账号对应用户ID
	TransInType  string `json:"trans_in_email,omitempty"`  // 转入人支付宝账号
}

type RefundQueryTradeResponseModel struct {
	Data RefundQueryTradeResponse `json:"alipay_trade_fastpay_refund_query_response"` // 返回值信息
	Sign string                   `json:"sign"`                                       // 签名，参见https://docs.open.alipay.com/291/106074
}
