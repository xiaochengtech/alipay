package alipay

import (
	"encoding/json"
)

// 统一收单线下交易查询接口
func (c *Client) QueryTrade(body QueryTradeBody) (aliRsp QueryTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.query", body, params, false)
	if err != nil {
		return
	}
	var response QueryTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	/* if err = c.verifySign(response.Data, response.Sign); err != nil {
		return
	} */
	aliRsp = response.Data
	return
}

type QueryTradeBody struct {
	OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 商户订单号
	TradeNo      string   `json:"trade_no,omitempty"`      // 支付宝交易号，和商户订单号不能同时为空
	OrgPid       string   `json:"org_pid,omitempty"`       // 交易所属收单机构的pid
	QueryOptions []string `json:"query_options,omitempty"` // 查询选项,定制查询返回信息
}

type QueryTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo          string              `json:"trade_no"`                         // 支付宝交易号
	OutTradeNo       string              `json:"out_trade_no"`                     // 商户订单号
	BuyerLogonId     string              `json:"buyer_logon_id"`                   // 买家支付宝账号
	TradeStatus      string              `json:"trade_status"`                     // 交易状态
	TotalAmount      string              `json:"total_amount"`                     // 订单总金额
	TransCurrency    string              `json:"trans_currency,omitempty"`         // 标价币种
	SettleCurrency   string              `json:"settle_currency,omitempty"`        // 订单结算币种
	SettleAmount     string              `json:"settle_amount,omitempty"`          // 结算币种订单金额
	PayCurrency      string              `json:"pay_currency,omitempty"`           // 订单支付币种
	PayAmount        string              `json:"pay_amount,omitempty"`             // 支付币种订单金额
	SettleTransRate  string              `json:"settle_trans_rate,omitempty"`      // 结算币种兑换标价币种汇率
	TransPayRate     string              `json:"trans_pay_rate,omitempty"`         // 标价币种兑换支付币种汇率
	BuyerPayAmount   string              `json:"buyer_pay_amount,omitempty"`       // 买家实付金额
	PointAmount      string              `json:"point_amount,omitempty"`           // 积分支付的金额
	InvoiceAmount    string              `json:"invoice_amount,omitempty"`         // 可开具发票的金额
	SendPayDate      string              `json:"send_pay_date,omitempty"`          // 本次交易打款给卖家的时间
	ReceiptAmount    string              `json:"receipt_amount,omitempty"`         // 实收金额
	StoreId          string              `json:"store_id,omitempty"`               // 商户门店编号
	TerminalId       string              `json:"terminal_id,omitempty"`            // 商户机具终端编号
	FundBillList     *[]FundBillListInfo `json:"fund_bill_list"`                   // 交易支付使用的资金渠道
	StoreName        string              `json:"store_name,omitempty"`             // 请求交易支付中的商户店铺的名称
	BuyerUserId      string              `json:"buyer_user_id"`                    // 买家在支付宝的用户id
	ChargeAmount     string              `json:"charge_amount,omitempty"`          // 该笔交易针对收款方的收费金额
	ChargeFlags      string              `json:"charge_flags,omitempty"`           // 费率活动标识
	SettlementId     string              `json:"settlement_id,omitempty"`          // 支付清算编号
	TradeSettle      *TradeSettleInfo    `json:"trade_settle_info,omitempty"`      // 返回的交易结算信息
	AuthTradePayMode string              `json:"auth_trade_pay_mode,omitempty"`    // 预授权支付模式
	BuyerUserType    string              `json:"buyer_user_type,omitempty"`        // 买家用户类型
	MdiscountAmount  string              `json:"mdiscount_amount,omitempty"`       // 商家优惠金额
	DiscountAmount   string              `json:"discount_amount,omitempty"`        // 平台优惠金额
	BuyerUserName    string              `json:"buyer_user_name,omitempty"`        // 买家为个人用户时为买家姓名，买家为企业用户时为企业名称
	Subject          string              `json:"subject,omitempty"`                // 订单标题
	Body             string              `json:"body,omitempty"`                   // 订单描述
	SubMerchantId    string              `json:"alipay_sub_merchant_id,omitempty"` // 间连商户在支付宝端的商户编号
	ExtInfos         string              `json:"ext_infos,omitempty"`              // 交易额外信息
}

type FundBillListInfo struct {
	FundChannel string `json:"fund_channel"`          // 交易使用的资金渠道
	BankCode    string `json:"bank_code,omitempty"`   // 银行卡支付时的银行代码
	Amount      string `json:"amount"`                // 该支付工具类型所使用的金额
	RealAmount  string `json:"real_amount,omitempty"` // 渠道实际付款金额
	FundType    string `json:"fund_type,omitempty"`   // 渠道所使用的资金类型
}

type TradeSettleInfo struct {
	TradeSettleDetailList []TradeSettleDetail `json:"trade_settle_detail_list,omitempty"` // 交易结算明细信息
}

type TradeSettleDetail struct {
	OperationType     string  `json:"operation_type"`                // 结算操作类型
	OperationSerialNo string  `json:"operation_serial_no,omitempty"` // 商户操作序列号
	OperationDate     string  `json:"operation_dt"`                  // 操作日期
	TransOut          string  `json:"trans_out,omitempty"`           // 转出账号
	TransIn           string  `json:"trans_in,omitempty"`            // 转入账号
	Amount            float32 `json:"amount"`                        // 实际操作金额
}

type QueryTradeResponseModel struct {
	Data QueryTradeResponse `json:"alipay_trade_query_response"` // 返回值信息
	Sign string             `json:"sign"`                        // 签名，参见https://docs.open.alipay.com/291/106074
}
