package alipay

import (
	"encoding/json"
	"time"
)

// 统一收单线下交易查询接口
func (c *Client) QueryTrade(body QueryTradeBody) (aliRsp QueryTradeResponse, err error) {
	params := BodyMap{}
	bytes, err := c.doAliPay("alipay.trade.query", body, params, true)
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
	OutTradeNo   string  `json:"out_trade_no,omitempty"`  // 商户订单号
	TradeNo      string  `json:"trade_no,omitempty"`      // 支付宝交易号，和商户订单号不能同时为空
	OrgPid       float32 `json:"org_pid,omitempty"`       // 交易所属收单机构的pid
	QueryOptions float32 `json:"query_options,omitempty"` // 查询选项,定制查询返回信息
}

type QueryTradeResponse struct {
	ResponseModel
	// 响应参数
	TradeNo             string           `json:"trade_no"`               // 支付宝交易号
	OutTradeNo          string           `json:"out_trade_no"`           // 商户订单号
	BuyerLogonId        string           `json:"buyer_logon_id"`         // 买家支付宝账号
	TradeStatus         string           `json:"trade_status"`           // 交易状态
	TotalAmount         string           `json:"total_amount"`           // 订单总金额
	TransCurrency       string           `json:"trans_currency"`         // 标价币种
	SettleCurrency      string           `json:"settle_currency"`        // 订单结算币种
	SettleAmount        string           `json:"settle_amount"`          // 结算币种订单金额
	PayCurrency         string           `json:"pay_currency"`           // 订单支付币种
	PayAmount           string           `json:"pay_amount"`             // 支付币种订单金额
	SettleTransRate     string           `json:"settle_trans_rate"`      // 结算币种兑换标价币种汇率
	TransPayRate        string           `json:"trans_pay_rate"`         // 标价币种兑换支付币种汇率
	BuyerPayAmount      string           `json:"buyer_pay_amount"`       // 买家实付金额
	PointAmount         string           `json:"point_amount"`           // 积分支付的金额
	InvoiceAmount       string           `json:"invoice_amount"`         // 可开具发票的金额
	SendPayDate         time.Time        `json:"send_pay_date"`          // 本次交易打款给卖家的时间
	ReceiptAmount       string           `json:"receipt_amount"`         // 实收金额
	StoreId             string           `json:"store_id"`               // 商户门店编号
	TerminalId          string           `json:"terminal_id"`            // 商户机具终端编号
	FundBillList        FundBillListInfo `json:"fund_bill_list"`         // 交易支付使用的资金渠道
	StoreName           string           `json:"store_name"`             // 请求交易支付中的商户店铺的名称
	BuyerUserId         string           `json:"buyer_user_id"`          // 买家在支付宝的用户id
	ChargeAmount        string           `json:"charge_amount"`          // 该笔交易针对收款方的收费金额
	ChargeFlags         string           `json:"charge_flags"`           // 费率活动标识
	SettlementId        string           `json:"settlement_id"`          // 支付清算编号
	TradeSettle         TradeSettleInfo  `json:"trade_settle_info"`      // 返回的交易结算信息
	AuthTradePayMode    string           `json:"auth_trade_pay_mode"`    // 预授权支付模式
	BuyerUserType       string           `json:"buyer_user_type"`        // 买家用户类型
	MdiscountAmount     string           `json:"mdiscount_amount"`       // 商家优惠金额
	DiscountAmount      string           `json:"discount_amount"`        // 平台优惠金额
	BuyerUserName       string           `json:"buyer_user_name"`        // 买家为个人用户时为买家姓名，买家为企业用户时为企业名称
	Subject             string           `json:"subject"`                // 订单标题
	Body                string           `json:"body"`                   // 订单描述
	AlipaySubMerchantId string           `json:"alipay_sub_merchant_id"` // 间连商户在支付宝端的商户编号
	ExtInfos            string           `json:"ext_infos"`              // 交易额外信息
}

type FundBillListInfo struct {
	FundChannel string  `json:"fund_channel"`          // 交易使用的资金渠道
	BankCode    string  `json:"bank_code,omitempty"`   // 银行卡支付时的银行代码
	Amount      float32 `json:"amount"`                // 该支付工具类型所使用的金额
	RealAmount  float32 `json:"real_amount,omitempty"` // 渠道实际付款金额
	FundType    string  `json:"fund_type,omitempty"`   // 渠道所使用的资金类型
}

type TradeSettleInfo struct {
	TradeSettleDetailList []TradeSettleDetail `json:"trade_settle_detail_list"` // 交易结算明细信息
}

type TradeSettleDetail struct {
	OperationType     string    `json:"operation_type"`      // 结算操作类型
	OperationSerialNo string    `json:"operation_serial_no"` // 商户操作序列号
	OperationDate     time.Time `json:"operation_dt"`        // 操作日期
	TransOut          string    `json:"trans_out"`           // 转出账号
	TransIn           string    `json:"trans_in"`            // 转入账号
	Amount            float32   `json:"amount"`              // 实际操作金额
}

type QueryTradeResponseModel struct {
	Data QueryTradeResponse `json:"alipay_trade_query_response"` // 返回值信息
	Sign string             `json:"sign"`                        // 签名，参见https://docs.open.alipay.com/291/106074
}
