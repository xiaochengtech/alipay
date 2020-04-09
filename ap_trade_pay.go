package alipay

import (
	"encoding/json"
)

// 统一收单交易支付接口
func (c Client) TradePay(body TradePayBody) (aliRsp TradePayResponse, err error) {
	params := BodyMap{
		"biz_content": c.GenerateBizContent(body),
	}
	bytes, err := c.doAlipay("alipay.trade.pay", params)
	if err != nil {
		return
	}
	var response TradePayResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type TradePayBody struct {
	OutTradeNo         string       `json:"out_trade_no"`                   // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	Scene              string       `json:"scene"`                          // 支付场景
	AuthCode           string       `json:"auth_code"`                      // 支付授权码
	ProductCode        string       `json:"product_code,omitempty"`         // 销售产品码
	Subject            string       `json:"subject"`                        // 订单标题
	BuyerId            string       `json:"buyer_id,omitempty"`             // 买家的支付宝唯一用户号
	SellerId           string       `json:"seller_id,omitempty"`            // 卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount        float32      `json:"total_amount,omitempty"`         // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	TransCurrency      string       `json:"trans_currency,omitempty"`       // 标价币种
	SettleCurrency     string       `json:"settle_currency,omitempty"`      // 订单结算币种
	DiscountAmount     float32      `json:"discountable_amount,omitempty"`  // 可打折金额
	Body               string       `json:"body,omitempty"`                 // 对交易或商品的描述
	GoodsDetail        *[]Goods     `json:"goods_detail,omitempty"`         // 订单包含的商品列表信息
	OperatorId         string       `json:"operator_id,omitempty"`          // 商户操作员编号
	StoreId            string       `json:"store_id,omitempty"`             // 商户门店编号
	TerminalId         string       `json:"terminal_id,omitempty"`          // 商户机具终端编号
	ExtendParams       *ExtendParam `json:"extend_params,omitempty"`        // 业务扩展参数
	TimeoutExpress     string       `json:"timeout_express,omitempty"`      // 该笔订单允许的最晚付款时间，逾期将关闭交易
	AuthConfirmMode    string       `json:"auth_confirm_mode,omitempty"`    // 预授权确认模式
	TerminalParams     string       `json:"terminal_params,omitempty"`      // 商户传入终端设备相关信息
	PromoParams        *PromoParam  `json:"promo_params,omitempty"`         // 优惠明细参数
	AdvancePaymentType string       `json:"advance_payment_type,omitempty"` // 支付模式类型
}

type PromoParam struct {
	ActualOrderTime string `json:"actual_order_time,omitempty"` // 存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间
}

type TradePayResponse struct {
	ResponseModel
	// 响应参数
	TradeNo             string              `json:"trade_no"`                        // 支付宝交易号
	OutTradeNo          string              `json:"out_trade_no"`                    // 商户订单号
	BuyerLogonId        string              `json:"buyer_logon_id"`                  // 买家支付宝账号
	SettleAmount        string              `json:"settle_amount,omitempty"`         // 结算币种订单金额
	PayCurrency         string              `json:"pay_currency,omitempty"`          // 订单支付币种
	PayAmount           string              `json:"pay_amount,omitempty"`            // 支付币种订单金额
	SettleTransRate     string              `json:"settle_trans_rate,omitempty"`     // 结算币种兑换标价币种汇率
	TransPayRate        string              `json:"trans_pay_rate,omitempty"`        // 标价币种兑换支付币种汇率
	TotalAmount         string              `json:"total_amount"`                    // 订单总金额
	TransCurrency       string              `json:"trans_currency,omitempty"`        // 标价币种
	SettleCurrency      string              `json:"settle_currency,omitempty"`       // 订单结算币种
	ReceiptAmount       string              `json:"receipt_amount"`                  // 实收金额
	BuyerPayAmount      string              `json:"buyer_pay_amount,omitempty"`      // 买家实付金额
	PointAmount         string              `json:"point_amount,omitempty"`          // 积分支付的金额
	InvoiceAmount       string              `json:"invoice_amount,omitempty"`        // 可开具发票的金额
	GmtPayment          string              `json:"gmt_payment"`                     // 交易支付时间
	FundBillList        *[]FundBillListInfo `json:"fund_bill_list"`                  // 交易支付使用的资金渠道
	CardBalance         string              `json:"card_balance,omitempty"`          // 支付宝卡余额
	StoreName           string              `json:"store_name,omitempty"`            // 请求交易支付中的商户店铺的名称
	BuyerUserId         string              `json:"buyer_user_id"`                   // 买家在支付宝的用户id
	DiscountGoodsDetail string              `json:"discount_goods_detail,omitempty"` // 本次交易支付所使用的单品券优惠的商品优惠信息
	VoucherDetailList   *[]VoucherDetail    `json:"voucher_detail_list,omitempty"`   // 本交易支付时使用的所有优惠券信息
	AdvanceAmount       string              `json:"advance_amount,omitempty"`        // 先享后付2.0垫资金额
	AuthTradePayMode    string              `json:"auth_trade_pay_mode,omitempty"`   // 预授权支付模式
	ChargeAmount        string              `json:"charge_amount,omitempty"`         // 该笔交易针对收款方的收费金额
	ChargeFlags         string              `json:"charge_flags,omitempty"`          // 费率活动标识
	SettlementId        string              `json:"settlement_id,omitempty"`         // 支付清算编号
	BusinessParams      string              `json:"business_params,omitempty"`       // 商户传入业务信息
	BuyerUserType       string              `json:"buyer_user_type,omitempty"`       // 买家用户类型
	MdiscountAmount     string              `json:"mdiscount_amount,omitempty"`      // 商家优惠金额
	DiscountAmount      string              `json:"discount_amount,omitempty"`       // 平台优惠金额
	BuyerUserName       string              `json:"buyer_user_name,omitempty"`       // 买家为个人用户时为买家姓名，买家为企业用户时为企业名称
}

type VoucherDetail struct {
	Id                         string `json:"id"`                                     // 券id
	Name                       string `json:"name"`                                   // 券名称
	Type                       string `json:"type"`                                   // 券类型
	Amount                     string `json:"amount"`                                 // 优惠券面额
	MerchantContribute         string `json:"merchant_contribute,omitempty"`          // 商家出资金额
	OtherContribute            string `json:"other_contribute,omitempty"`             // 其他出资方出资金额
	Memo                       string `json:"memo,omitempty"`                         // 优惠券备注信息
	TemplateId                 string `json:"template_id,omitempty"`                  // 券模板id
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`    // 用户实际付款的金额
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"` // 商户优惠的金额
	PurchaseMAntContribute     string `json:"purchase_ant_contribute,omitempty"`      // 平台优惠的金额
}

type TradePayResponseModel struct {
	Data TradePayResponse `json:"alipay_trade_pay_response"` // 返回值信息
	Sign string           `json:"sign"`                      // 签名，参见https://docs.open.alipay.com/291/106074
}
