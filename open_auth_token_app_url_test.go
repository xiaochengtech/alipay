package alipay

import (
	"fmt"
	"testing"
)

func TestGetOpenAuthTokenAppURL(t *testing.T) {
	fmt.Println("----------第三方应用授权URL----------")
	appId, redirectURI := "2016102200736537", "https://www.baidu.com"
	url := GetOpenAuthTokenAppURL(false, appId, redirectURI)
	fmt.Printf("返回值: %+v\n", url)
	return
}
