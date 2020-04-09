package alipay

import (
	"fmt"
	"net/url"
)

// 第三方应用授权URL
// https://opendocs.alipay.com/open/20160728150111277227/intro
func GetOpenAuthTokenAppURL(isProd bool, appId string, redirectUri string) (result string) {
	if isProd {
		result += "https://openauth.alipay.com"
	} else {
		result += "https://openauth.alipaydev.com"
	}
	result += "/oauth2/appToAppAuth.htm"
	result += fmt.Sprintf("?app_id=%s", appId)
	result += fmt.Sprintf("&redirect_uri=%s", url.QueryEscape(redirectUri))
	return
}
