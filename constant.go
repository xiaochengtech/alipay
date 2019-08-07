package alipay

type BodyMap map[string]interface{}

const (
	baseUrl        = "https://openapi.alipay.com/gateway.do"    // (生产环境) 支付宝接口地址
	baseUrlSandbox = "https://openapi.alipaydev.com/gateway.do" // (沙盒环境) 支付宝接口地址

	// 请求格式
	FormatJson = "JSON" // Json格式

	// 签名算法类型
	SignTypeRSA  = "RSA"  // RSA签名
	SignTypeRSA2 = "RSA2" // RSA2签名

	// 版本号
	Version1 = "1.0" // 1.0版本
)
