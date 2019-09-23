package alipay

import (
	"encoding/json"
)

// 统一收单交易预创接口，支付动作在支付宝内完成，根据传入的通知地址异步通知服务商（推荐），
func (c *Client) PreCreateTrade(body PreCreateTradeBody, notifyUrl string) (aliRsp PreCreateTradeResponse, err error) {
	params := BodyMap{}
	// 按需设置公共请求参数
	if len(notifyUrl) > 0 {
		params["notify_url"] = notifyUrl
	}
	bytes, err := c.doAliPay("alipay.trade.precreate", body, params, false)
	if err != nil {
		return
	}
	var response PreCreateTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type PreCreateTradeBody struct {
	OutTradeNo           string         `json:"out_trade_no"`                      // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	SellerId             string         `json:"seller_id,omitempty"`               // 卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount          float32        `json:"total_amount"`                      // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	DiscountAmount       float32        `json:"discountable_amount,omitempty"`     // 可打折金额
	Subject              string         `json:"subject"`                           // 订单标题
	GoodsDetail          []Goods        `json:"goods_detail,omitempty"`            // 订单包含的商品列表信息
	Body                 string         `json:"body,omitempty"`                    // 对交易或商品的描述
	ProductCode          string         `json:"product_code,omitempty"`            // 销售产品码
	OperatorId           string         `json:"operator_id,omitempty"`             // 商户操作员编号
	StoreId              string         `json:"store_id,omitempty"`                // 商户门店编号
	DisablePayChannels   string         `json:"disable_pay_channels,omitempty"`    // 禁用渠道
	EnablePayChannels    string         `json:"enable_pay_channels,omitempty"`     // 可用渠道
	TerminalId           string         `json:"terminal_id,omitempty"`             // 商户机具终端编号
	ExtendParams         *ExtendParam   `json:"extend_params,omitempty"`           // 业务扩展参数
	TimeoutExpress       string         `json:"timeout_express,omitempty"`         // 该笔订单允许的最晚付款时间，逾期将关闭交易
	SettleInfo           *Settle        `json:"settle_info,omitempty"`             // 描述结算信息
	MerchantOrderNo      string         `json:"merchant_order_no,omitempty"`       // 商户原始订单号
	BusinessParams       *BusinessParam `json:"business_params,omitempty"`         // 商户传入业务信息
	QrCodeTimeoutExpress string         `json:"qr_code_timeout_express,omitempty"` // 该笔订单允许的最晚付款时间
}

type PreCreateTradeResponse struct {
	ResponseModel
	// 响应参数
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	QrCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串
}

type PreCreateTradeResponseModel struct {
	Data PreCreateTradeResponse `json:"alipay_trade_precreate_response"` // 返回值信息
	Sign string                 `json:"sign"`                            // 签名，参见https://docs.open.alipay.com/291/106074
}
