package alipay

// 公共响应参数
type ResponseModel struct {
	Code    string `json:"code"`     // 网关返回码，参见https://docs.open.alipay.com/common/105806
	Msg     string `json:"msg"`      // 网关返回码描述，参见https://docs.open.alipay.com/common/105806
	SubCode string `json:"sub_code"` // 业务返回码，参见具体的API接口文档
	SubMsg  string `json:"sub_msg"`  // 业务返回码描述，参见具体的API接口文档
	Sign    string `json:"sign"`     // 签名，参见https://docs.open.alipay.com/291/106074
}
