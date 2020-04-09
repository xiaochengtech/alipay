package alipay

import (
	"encoding/json"
)

// 换取应用授权令牌
// https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
func (c *Client) OpenAuthTokenApp(body OpenAuthTokenAppBody) (aliRsp OpenAuthTokenAppResponse, err error) {
	params := c.ConvertToBodyMap(body)
	bytes, err := c.doAlipay("alipay.open.auth.token.app", params)
	if err != nil {
		return
	}
	var response OpenAuthTokenAppResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.Data
	return
}

type OpenAuthTokenAppBody struct {
	GrantType    string `json:"grant_type"`              // 参见constant.go。值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code,omitempty"`          // 授权码，如果grant_type的值为authorization_code，该值必须填写
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新令牌，如果grant_type值为refresh_token，该值不能为空
}

type OpenAuthTokenAppResponse struct {
	ResponseModel
	UserId          string `json:"user_id"`           // 授权商户的user_id
	AuthAppId       string `json:"auth_app_id"`       // 授权商户的appid
	AppAuthToken    string `json:"app_auth_token"`    // 应用授权令牌
	AppRefreshToken string `json:"app_refresh_token"` // 刷新令牌
	ExpiresIn       int64  `json:"expires_in"`        // 应用授权令牌的有效时间(从接口调用时间作为起始时间)，单位到秒
	ReExpiresIn     int64  `json:"re_expires_in"`     // 刷新令牌的有效时间(从接口调用时间作为起始时间)，单位到秒
}

type OpenAuthTokenAppResponseModel struct {
	Data OpenAuthTokenAppResponse `json:"alipay_open_auth_token_app_response"` // 返回值信息
	Sign string                   `json:"sign"`                                // 签名，参见https://docs.open.alipay.com/291/106074
}
