package alipay

import (
	"encoding/json"
)

// 换取授权访问令牌接口
// https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (c Client) SystemOauthToken(body SystemOauthTokenBody) (aliRsp SystemOauthTokenResponse, err error) {
	params := c.ConvertToBodyMap(body)
	bytes, err := c.doAlipay("alipay.system.oauth.token", params)
	if err != nil {
		return
	}
	var response SystemOauthTokenResponseModel
	if err = json.Unmarshal(bytes, &response); err != nil {
		return
	}
	aliRsp = response.SystemOauthTokenResponse
	return
}

type SystemOauthTokenBody struct {
	GrantType    string `json:"grant_type"`              // 参见constant.go。值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code,omitempty"`          // 授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
}

type SystemOauthTokenData struct {
	UserId       string `json:"user_id"`       // 支付宝用户的唯一userId，2088102150477652
	AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口，20120823ac6ffaa4d2d84e7384bf983531473993
	ExpiresIn    int64  `json:"expires_in"`    // 访问令牌的有效时间，单位是秒。3600
	RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token，20120823ac6ffdsdf2d84e7384bf983531473993
	ReExpiresIn  int64  `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。3600
}

type SystemOauthTokenResponse struct {
	Error ResponseModel        `json:"error_response,omitempty"`
	Data  SystemOauthTokenData `json:"alipay_system_oauth_token_response"` // 返回值信息
}

type SystemOauthTokenResponseModel struct {
	SystemOauthTokenResponse
	Sign string `json:"sign"` // 签名，参见https://docs.open.alipay.com/291/106074
}
