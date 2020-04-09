package alipay

import (
	"encoding/json"
)

// 统一收单交易退款接口
func (c Client) TradeRefund(body TradeRefundBody) (aliRsp TradeRefundResponse, err error) {
	params := BodyMap{
		"biz_content": c.GenerateBizContent(body),
	}
	bytes, err := c.doAlipay("alipay.trade.refund", params)
	if err != nil {
		return
	}
	var response TradeRefundResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type TradeRefundBody struct {
	OutTradeNo              string             `json:"out_trade_no,omitempty"`              // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	TradeNo                 string             `json:"trade_no,omitempty"`                  // 支付宝订单号
	RefundAmount            float32            `json:"refund_amount"`                       // 退款金额
	RefundCurrency          string             `json:"refund_currency,omitempty"`           // 退款币种
	RefundReason            string             `json:"refund_reason,omitempty"`             // 退款的原因说明
	OutRequestNo            string             `json:"out_request_no,omitempty"`            // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传
	OperatorId              string             `json:"operator_id,omitempty"`               // 商户操作员编号
	StoreId                 string             `json:"store_id,omitempty"`                  // 商户门店编号
	TerminalId              string             `json:"terminal_id,omitempty"`               // 商户机具终端编号
	GoodsDetail             []Goods            `json:"goods_detail,omitempty"`              // 退款包含的商品列表信息
	RefundRoyaltyParameters []RoyaltyParameter `json:"refund_royalty_parameters,omitempty"` // 退分账明细信息
	OrgPId                  string             `json:"org_pid,omitempty"`                   // 退款的交易所属收单机构的pid
}

type RoyaltyParameter struct {
	RoyaltyType      string  `json:"royalty_type,omitempty"`      // 分账类型
	TransOut         string  `json:"trans_out,omitempty"`         // 支出方账户
	TransOutType     string  `json:"trans_out_type,omitempty"`    // 支出方账户类型
	TransInType      string  `json:"trans_in_type,omitempty"`     // 收入方账户类型
	TransIn          string  `json:"trans_in"`                    // 分账类型
	Amount           float32 `json:"amount,omitempty"`            // 分账的金额
	AmountPercentage int     `json:"amount_percentage,omitempty"` // 分账信息中分账百分比,取值范围为大于0，少于或等于100的整数。
	Desc             string  `json:"desc,omitempty"`              // 分账描述
}

type TradeRefundResponse struct {
	ResponseModel
	// 响应参数
	TradeNo               string             `json:"trade_no"`                                  // 支付宝交易号
	OutTradeNo            string             `json:"out_trade_no"`                              // 商户订单号
	BuyerLogonId          string             `json:"buyer_logon_id"`                            // 买家支付宝账号
	FundChange            string             `json:"fund_change"`                               // 本次退款是否发生了资金变化
	RefundFee             string             `json:"refund_fee"`                                // 退款总金额
	RefundCurrency        string             `json:"refund_currency,omitempty"`                 // 退款币种信息
	GmtRefundPay          string             `json:"gmt_refund_pay"`                            // 退款支付时间
	FundBillList          []FundBillListInfo `json:"refund_detail_item_list,omitempty"`         // 退款使用的资金渠道
	StoreName             string             `json:"store_name,omitempty"`                      // 请求交易支付中的商户店铺的名称
	BuyerUserId           string             `json:"buyer_user_id"`                             // 买家在支付宝的用户id
	RefundPaytoolList     *PresetPayToolInfo `json:"refund_preset_paytool_list,omitempty"`      // 退回的前置资产列表
	RefundSettlementId    string             `json:"refund_settlement_id,omitempty"`            // 支付清算编号
	RefundBuyerAmount     string             `json:"present_refund_buyer_amount,omitempty"`     // 本次退款金额中买家退款金额
	RefundDiscountAmount  string             `json:"present_refund_discount_amount,omitempty"`  // 本次退款金额中平台优惠退款金额
	RefundMdiscountAmount string             `json:"present_refund_mdiscount_amount,omitempty"` // 本次退款金额中商家优惠退款金额
}

type PresetPayToolInfo struct {
	Amount         []string `json:"amount"`           // 前置资产金额
	AssertTypeCode string   `json:"assert_type_code"` // 前置资产类型编码

}

type TradeRefundResponseModel struct {
	Data TradeRefundResponse `json:"alipay_trade_refund_response"` // 返回值信息
	Sign string              `json:"sign"`                         // 签名，参见https://docs.open.alipay.com/291/106074
}
