package alipay

import (
	"encoding/json"
)

// 统一收单交易创建接口，对应pay接口支付
func (c *Client) CreateTrade(body CreateTradeBody, notifyUrl string) (aliRsp CreateTradeResponse, err error) {
	params := BodyMap{}
	// 按需设置公共请求参数
	if len(notifyUrl) > 0 {
		params["notify_url"] = notifyUrl
	}
	bytes, err := c.doAlipay("alipay.trade.create", body, params, false, true)
	if err != nil {
		return
	}
	var response CreateTradeResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type CreateTradeBody struct {
	OutTradeNo          string           `json:"out_trade_no"`                    // 商户订单号，64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	SellerId            string           `json:"seller_id,omitempty"`             // 卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount         float32          `json:"total_amount"`                    // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	DiscountAmount      float32          `json:"discountable_amount,omitempty"`   // 可打折金额
	Subject             string           `json:"subject"`                         // 订单标题
	Body                string           `json:"body,omitempty"`                  // 对交易或商品的描述
	BuyerId             string           `json:"buyer_id,omitempty"`              // 买家的支付宝唯一用户号
	GoodsDetail         []Goods          `json:"goods_detail,omitempty"`          // 订单包含的商品列表信息
	ProductCode         string           `json:"product_code,omitempty"`          // 销售产品码
	OperatorId          string           `json:"operator_id,omitempty"`           // 商户操作员编号
	StoreId             string           `json:"store_id,omitempty"`              // 商户门店编号
	TerminalId          string           `json:"terminal_id,omitempty"`           // 商户机具终端编号
	ExtendParams        *ExtendParam     `json:"extend_params,omitempty"`         // 业务扩展参数
	TimeoutExpress      string           `json:"timeout_express,omitempty"`       // 该笔订单允许的最晚付款时间，逾期将关闭交易
	SettleInfo          *Settle          `json:"settle_info,omitempty"`           // 描述结算信息
	LogisticsDetail     *Logistics       `json:"logistics_detail,omitempty"`      // 物流信息
	BusinessParams      *BusinessParam   `json:"business_params,omitempty"`       // 商户传入业务信息
	ReceiverAddressInfo *ReceiverAddress `json:"receiver_address_info,omitempty"` // 收货人及地址信息
}

type Goods struct {
	GoodsId        string  `json:"goods_id"`                  // 商品的编号
	GoodsName      string  `json:"goods_name"`                // 商品名称
	Quantity       int     `json:"quantity"`                  // 商品数量
	Price          float32 `json:"price"`                     // 商品单价，单位为元
	GoodsCategory  string  `json:"goods_category,omitempty"`  // 商品类目
	CategoriesTree string  `json:"categories_tree,omitempty"` // 商品类目树
	Body           string  `json:"body,omitempty"`            // 商品描述信息
	ShowUrl        string  `json:"show_url,omitempty"`        // 商品的展示地址
}

type ExtendParam struct {
	ProviderId string `json:"sys_service_provider_id,omitempty"` // 系统商编号
	Reflux     string `json:"industry_reflux_info,omitempty"`    // 行业数据回流信息
	CardType   string `json:"card_type,omitempty"`               // 卡类型
}

type Settle struct {
	SettleDetailInfos []SettleDetail `json:"settle_detail_infos"`     // 结算详细信息
	MerchantType      string         `json:"merchant_type,omitempty"` // 商户id类型
}

type SettleDetail struct {
	TransInType      string  `json:"trans_in_type"`                // 结算收款方的账户类型(见constant定义)
	TransIn          string  `json:"trans_in"`                     // 结算收款方
	SummaryDimension string  `json:"summary_dimension,omitempty"`  // 结算汇总维度
	SettleEntityId   string  `json:"settle_entity_id,omitempty"`   // 结算主体标识
	SettleEntityType string  `json:"settle_entity_type,omitempty"` // 结算主体类型
	Amount           float32 `json:"amount"`                       // 结算的金额，单位为元。目前必须和交易金额相同
}

type Logistics struct {
	LogisticsType string `json:"logistics_type,omitempty"` // 物流类型(见constant定义)
}

type BusinessParam struct {
	CampusCard      string `json:"campus_card,omitempty"`       // 校园卡编号
	CardType        string `json:"card_type,omitempty"`         // 虚拟卡卡类型
	ActualOrderTime string `json:"actual_order_time,omitempty"` // 实际订单时间
}

type ReceiverAddress struct {
	Name         string `json:"name,omitempty"`          // 收货人的姓名
	Address      string `json:"address,omitempty"`       // 收货地址
	Mobile       string `json:"mobile,omitempty"`        // 收货人手机号
	Zip          string `json:"zip,omitempty"`           // 收货地址邮编
	DivisionCode string `json:"division_code,omitempty"` // 中国标准城市区域码
}

type CreateTradeResponse struct {
	ResponseModel
	// 响应参数
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
}

type CreateTradeResponseModel struct {
	Data CreateTradeResponse `json:"alipay_trade_create_response"` // 返回值信息
	Sign string              `json:"sign"`                         // 签名，参见https://docs.open.alipay.com/291/106074
}
