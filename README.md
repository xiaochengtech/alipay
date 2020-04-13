# alipay

这是用Golang封装了支付宝的API接口的SDK。

* 支持服务商或者普通商户模式。
* 全部参数和返回值均使用`struct`类型传递，而不是`map`类型。

### 初始化

```go
const (
	isProd = true
)
config := alipay.Config{
	AppId: "xxxxxx",
	AppAuthToken: "yyyyyy",
}
client := alipay.NewClient(isProd, publicKey, privateKey, config)
```

### 使用

对于非Client接口，可以直接调用，对于Client接口，需要先通过初始化生成client，然后调用相应方法：

```go
func Test() {
	// 初始化参数
	body := alipay.TradeCreateBody{}
	body.OutTradeNo = "GYWX201910311240354444"
	body.SellerId = "2088102178986262"
	body.TotalAmount = 2.00
	body.Subject = "测试车场阿里支付-停车费"
	body.BuyerId = "2088102179285843"
	notifyUrl := "http://www.example.com"
	// 请求支付
	aliRsp, err := client.TradeCreate(body, notifyUrl)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", aliRsp)
}
```

注意事项：

* 参数或返回值的类型，请查看接口对应的文件，里面有`XXXBody`和`XXXResponse`与之对应。
* 参数或返回值中的常量，请参照[constant.go](constant.go)文件。
* 具体使用方法，请参照接口对应的测试文件。

### 接口列表

对应实现文件`ap_xxx.go`，测试文件`ap_xxx_test.go`。

- [ ] 支付API
	- [x] 统一收单交易支付接口：`(Client) TradePay`，交易支付触发通知处理方法`(Client) NotifyPay`。
	- [x] 统一收单线下交易预创建：`(Client) TradePrecreate`。
	- [ ] app支付接口2.0
	- [ ] 手机网站支付接口2.0
	- [ ] 统一收单下单并支付页面接口
	- [x] 统一收单交易创建接口：`(Client) TradeCreate`。
	- [x] 统一收单线下交易查询：`(Client) TradeQuery`。
	- [x] 统一收单交易撤销接口：`(Client) TradeCancel`。
	- [x] 统一收单交易关闭接口：`(Client) TradeClose`。
	- [x] 统一收单交易退款接口：`(Client) TradeRefund`。
	- [ ] 统一收单退款页面接口
	- [x] 统一收单交易退款查询：`(Client) TradeFastpayRefundQuery`。
	- [x] 统一收单交易结算接口：`(Client) TradeOrderSettle`。
	- [ ] 资金授权冻结接口
	- [x] 支付宝订单信息同步接口：`(Client) TradeOrderinfoSync`。
	- [ ] 订单咨询服务
	- [ ] 聚合支付订单咨询服务
	- [ ] 花呗先享会员结算申请
	- [ ] NFC用户卡信息同步
	- [ ] 广告投放数据查询
	- [ ] 航司电话订票待申请接口
	- [ ] 网商银行全渠道收单业务订单创建
	- [ ] 口碑订单预下单
	- [ ] 口碑商品交易购买接口
	- [ ] 口碑订单预咨询
	- [ ] 口碑商品交易退货接口
	- [ ] 口碑商品交易查询接口
	- [ ] 码商发码成功回调接口
	- [ ] 口碑凭证延期接口
	- [ ] 口碑凭证码查询
	- [ ] 口碑凭证码撤销核销
	- [ ] 统一收单交易退款接口
- [ ] 会员API
- [ ] 店铺API
- [ ] 营销API
- [ ] 生活号API
- [ ] 芝麻信用API
- [ ] 工具类API
	- [ ] 用户登陆授权
	- [x] 换取授权访问令牌：`(Client) SystemOauthToken`。
	- [x] 换取应用授权令牌：`(Client) OpenAuthTokenApp`，授权URL生成方法`GetOpenAuthTokenAppURL`。
	- [ ] 查询某个应用授权AppAuthToken的授权信息
	- [ ] 应用支付宝公钥证书下载
	- [ ] 验签接口
	- [ ] 订阅消息主题
	- [ ] 变更订阅关系属性
	- [ ] 查询消息订阅关系
	- [ ] 取消消息订阅关系
	- [ ] 上报线下服务异常
	- [ ] 口碑业务授权令牌查询
- [ ] 风险控制API
- [ ] 服务市场API
- [ ] 账务API
- [ ] 生活缴费API
- [ ] 车主服务API
- [ ] 数据服务API
- [ ] 教育服务API
- [ ] 卡券API
- [ ] 广告API
- [ ] 资金API
- [ ] 地铁购票API
- [ ] 电子发票API
- [ ] 理财API
- [ ] 开放生态API
- [ ] 小程序API
- [ ] 历史API

### 文档

* 开放平台文档-API清单：[https://opendocs.alipay.com/apis](https://opendocs.alipay.com/apis)
* 网页&移动应用-第三方应用授权：[https://opendocs.alipay.com/open/20160728150111277227/intro](https://opendocs.alipay.com/open/20160728150111277227/intro)
* 沙箱账号：[https://openhome.alipay.com/platform/appDaily.htm](https://openhome.alipay.com/platform/appDaily.htm)
