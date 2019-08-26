# alipay

[![Latest Tag](https://img.shields.io/github/tag/cuckoopark/alipay.svg)](https://github.com/cuckoopark/alipay/releases/latest)

### 安装

```shell
go get -u github.com/cuckoopark/alipay
```

## 支付宝支付

* 手机网站支付：client.AliPayTradeWapPay()
* APP支付：client.AliPayTradeAppPay()
* 统一收单交易支付接口(当面付)：client.AliPayTradePay()
* 统一收单交易创建接口：client.AliPayTradeCreate()
* 统一收单线下交易查询：client.AliPayTradeQuery()
* 统一收单交易关闭接口：client.AliPayTradeClose()
* 统一收单交易撤销接口：client.AliPayTradeCancel()

## 支付宝公共API

* gopay.FormatPrivateKey() => 格式化应用私钥
* gopay.FormatAliPayPublicKey() => 格式化支付宝公钥
* gopay.ParseAliPayNotifyResult() => 解析并返回支付宝支付异步通知的参数
* gopay.VerifyAliPayResultSign() => 支付宝支付异步通知的签名验证和返回参数验签后的Sign

# 支付宝支付

<font color='#0088ff'>注意：具体请求参数根据请求的不同而不同，请参考支付宝官方文档的参数说明！</font>

支付宝官方文档：[官方文档](https://docs.open.alipay.com/catalog)

支付宝RSA秘钥生成文档：[生成 RSA 密钥](https://docs.open.alipay.com/291/105971/)

支付宝在线调试：[在线调试地址](https://openhome.alipay.com/platform/demoManage.htm)

沙箱环境使用说明：[文档地址](https://docs.open.alipay.com/200/105311)


### 1、支付结果异步通知参数解析；2、验签操作

> 支付宝支付后的异步通知验签文档[支付结果通知](https://docs.open.alipay.com/200/106120)

```go
//解析支付完成后的异步通知参数信息
//此处 c.Request() 为 *http.Request
notifyRsp, err := gopay.ParseAliPayNotifyResult(c.Request())
if err != nil {
    fmt.Println("gopay.ParseAliPayNotifyResult:", err)
    return
}
fmt.Println("notifyRsp:", notifyRsp)

aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
//验签操作
//    aliPayPublicKey：支付宝公钥
//    notifyRsp：利用 gopay.ParseAliPayNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数err：错误信息
ok, err := gopay.VerifyAliPayResultSign(aliPayPublicKey, notifyRsp)
if err != nil {
	log.Println("gopay.VerifyAliPayResultSign:", err)
	return
}
fmt.Println("ok:", ok)
```

### 支付宝付款结果异步通知,需回复支付宝平台是否成功

* 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

> 代码中return写法，由于本人用的[Echo Web框架](https://github.com/labstack/echo)，有兴趣的可以尝试一下

```go
return c.String(http.StatusOK, "success")
```

### 手机网站支付

* 手机网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面

> 文档说明[手机网站支付-请求参数说明](https://docs.open.alipay.com/203/107090/) 

> 文档说明[手机网站支付接口2.0](https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用秘钥
//    isProd：是否是正式环境
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//配置公共参数
client.SetCharset("UTF-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//请求参数
body := make(gopay.BodyMap)
body.Set("subject", "测试支付")
body.Set("out_trade_no", "GYWX201901301040355706100409")
body.Set("quit_url", "https://www.gopay.ink")
body.Set("total_amount", "10.00")
body.Set("product_code", "QUICK_WAP_WAY")
//手机网站支付请求
payUrl, err := client.AliPayTradeWapPay(body)
if err != nil {
	log.Println("err:", err)
	return
}
fmt.Println("payUrl:", payUrl)
```

### APP支付

* APP支付是通过服务端获取支付参数后，然后通过Android/iOS客户端的SDK调用支付功能

> 文档说明[APP支付-请求参数说明](https://docs.open.alipay.com/204/105465/) 

> 文档说明[APP支付接口2.0](https://docs.open.alipay.com/api_1/alipay.trade.app.pay/) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用秘钥
//    isProd：是否是正式环境
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//配置公共参数
client.SetCharset("UTF-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//请求参数
body := make(gopay.BodyMap)
body.Set("subject", "测试APP支付")
body.Set("out_trade_no", "GYWX201901301040355706100411")
body.Set("total_amount", "1.00")
//手机APP支付参数请求
payParam, err := client.AliPayTradeAppPay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("payParam:", payParam)
```

### 电脑网站支付

* 电脑网站支付是通过服务端获取支付URL后，然后返回给客户端，请求URL地址即可打开支付页面

> 文档说明[电脑网站支付](https://docs.open.alipay.com/270) 

> 文档说明[统一收单下单并支付页面接口](https://docs.open.alipay.com/api_1/alipay.trade.page.pay) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用秘钥
//    isProd：是否是正式环境
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//配置公共参数
client.SetCharset("UTF-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//请求参数
body := make(gopay.BodyMap)
body.Set("subject", "网站测试支付")
body.Set("out_trade_no", "GYWX201901301040355706100418")
body.Set("total_amount", "88.88")
body.Set("product_code", "FAST_INSTANT_TRADE_PAY")

//电脑网站支付请求
payUrl, err := client.AliPayTradePagePay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("payUrl:", payUrl)
```

### 当面付-条码支付

* 商家使用扫码枪等条码识别设备扫描用户支付宝钱包上的条码/二维码，完成收款。

> 文档说明[当面付-条码支付](https://docs.open.alipay.com/194) 

> 文档说明[统一收单交易支付接口](https://docs.open.alipay.com/api_1/alipay.trade.pay) 

```go
privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
//初始化支付宝客户端
//    appId：应用ID
//    privateKey：应用秘钥
//    isProd：是否是正式环境
client := gopay.NewAliPayClient("2016091200494382", privateKey, false)
//配置公共参数
client.SetCharset("UTF-8").
	SetSignType("RSA2").
	//SetAppAuthToken("").
	//SetReturnUrl("https://www.gopay.ink").
	SetNotifyUrl("https://www.gopay.ink")
//请求参数
body := make(gopay.BodyMap)
body.Set("subject", "条码支付")
body.Set("scene", "bar_code")
body.Set("auth_code", "285860185283886370")
body.Set("out_trade_no", "GYWX201901301040355706100456")
body.Set("total_amount", "10.00")
body.Set("timeout_express", "2m")

//当面付-条码支付
aliRsp, err := client.AliPayTradePay(body)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("aliRsp:", *aliRsp)
```