package alipay

type Client struct {
	config     Config // 配置信息
	publicKey  string // 支付宝公钥
	privateKey string // 应用私钥
	isProd     bool   // 是否是生产环境
}

// 初始化支付宝客户端
func NewClient(isProd bool, publicKey string, privateKey string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.publicKey = publicKey
	client.privateKey = privateKey
	client.isProd = isProd
	return client
}

//  测试用客户端
var testClient = NewClient(false, aliPayPublicKey, aliPayPrivateKey, Config{
	AppId:    "2016101000655815",
	Format:   FormatJson,
	Charset:  CharSetUTF8,
	SignType: SignTypeRSA2,
	Version:  Version1,
})

/*
//alipay.trade.fastpay.refund.query(统一收单交易退款查询)
func (this *aliPayClient) AliPayTradeFastPayRefundQuery(body BodyMap) {

}

//alipay.trade.order.settle(统一收单交易结算接口)
func (this *aliPayClient) AliPayTradeOrderSettle(body BodyMap) {

}


//alipay.trade.refund(统一收单交易退款接口)
func (this *aliPayClient) AliPayTradeRefund(body BodyMap) {

}

//alipay.trade.precreate(统一收单线下交易预创建)
func (this *aliPayClient) AliPayTradePrecreate(body BodyMap) {

}

//alipay.trade.page.pay(统一收单下单并支付页面接口)
func (this *aliPayClient) AliPayTradePagePay(body BodyMap) (payUrl string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	bytes, err = this.doAliPay(body, "alipay.trade.page.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	if payUrl == zfb_sanbox_base_url || payUrl == zfb_base_url {
		return null, errors.New("请求失败，请查看文档并检查参数")
	}
	return payUrl, nil
}

//alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
func (this *aliPayClient) AliPayTradeOrderinfoSync(body BodyMap) {

}
*/
