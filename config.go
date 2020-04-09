package alipay

// 公共参数配置
type Config struct {
	AppId        string `json:"app_id"`         // 支付宝分配给开发者的应用ID
	Format       string `json:"format"`         // (可不设置) 仅支持JSON
	Charset      string `json:"charset"`        // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType     string `json:"sign_type"`      // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Version      string `json:"version"`        // (可不设置) 调用的接口版本，固定为：1.0
	AppAuthToken string `json:"app_auth_token"` // 应用授权，参见https://docs.open.alipay.com/common/105193
}
