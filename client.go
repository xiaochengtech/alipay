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

//alipay.trade.close(统一收单交易关闭接口)
func (this *aliPayClient) AliPayTradeClose(body BodyMap) (aliRsp *AliPayTradeCloseResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.close")
	if err != nil {
		return nil, err
	}
	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("AliPayTradeClose::::", string(convertBytes))
	aliRsp = new(AliPayTradeCloseResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradeCloseResponse.Code != "10000" {
		info := aliRsp.AlipayTradeCloseResponse
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.cancel(统一收单交易撤销接口)
func (this *aliPayClient) AliPayTradeCancel(body BodyMap) (aliRsp *AliPayTradeCancelResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.cancel")
	if err != nil {
		return nil, err
	}
	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("AliPayTradeCancel::::", string(convertBytes))
	aliRsp = new(AliPayTradeCancelResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeCancelResponse.Code != "10000" {
		info := aliRsp.AliPayTradeCancelResponse
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.refund(统一收单交易退款接口)
func (this *aliPayClient) AliPayTradeRefund(body BodyMap) {

}

//alipay.trade.precreate(统一收单线下交易预创建)
func (this *aliPayClient) AliPayTradePrecreate(body BodyMap) {

}

//alipay.trade.pay(统一收单交易支付接口)
func (this *aliPayClient) AliPayTradePay(body BodyMap) (aliRsp *AliPayTradePayResponse, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return nil, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	//body.Set("product_code", "FACE_TO_FACE_PAYMENT")
	bytes, err = this.doAliPay(body, "alipay.trade.pay")
	if err != nil {
		return nil, err
	}

	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("convertBytes::::", string(convertBytes))
	aliRsp = new(AliPayTradePayResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradePayResponse.Code != "10000" {
		info := aliRsp.AlipayTradePayResponse
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.query(统一收单线下交易查询)
func (this *aliPayClient) AliPayTradeQuery(body BodyMap) (aliRsp *AliPayTradeQueryResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.query")
	if err != nil {
		return nil, err
	}
	convertBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
	//log.Println("AliPayTradeQuery::::", string(convertBytes))
	aliRsp = new(AliPayTradeQueryResponse)
	err = json.Unmarshal(convertBytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradeQueryResponse.Code != "10000" {
		info := aliRsp.AlipayTradeQueryResponse
		return nil, fmt.Errorf("code:%v,msg:%v,sub_code:%v,sub_msg:%v.", info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	return aliRsp, nil
}

//alipay.trade.app.pay(app支付接口2.0)
func (this *aliPayClient) AliPayTradeAppPay(body BodyMap) (payParam string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	//body.Set("product_code", "QUICK_MSECURITY_PAY")
	bytes, err = this.doAliPay(body, "alipay.trade.app.pay")
	if err != nil {
		return null, err
	}
	payParam = string(bytes)
	return payParam, nil
}

//alipay.trade.wap.pay(手机网站支付接口2.0)
func (this *aliPayClient) AliPayTradeWapPay(body BodyMap) (payUrl string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	body.Set("product_code", "QUICK_WAP_WAY")
	bytes, err = this.doAliPay(body, "alipay.trade.wap.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	//fmt.Println("URL::", payUrl)
	if payUrl == zfb_sanbox_base_url || payUrl == zfb_base_url {
		return null, errors.New("请求失败，请查看文档并检查参数")
	}
	return payUrl, nil
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
