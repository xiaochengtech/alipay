package alipay

import (
	"fmt"
	"testing"
)

func TestGetOpenAuthTokenAppURL(t *testing.T) {
	fmt.Println("----------第三方应用授权URL----------")
	appId, redirectURI := "2015101400446982", "http://example.com"
	url := GetOpenAuthTokenAppURL(true, appId, redirectURI)
	fmt.Printf("返回值: %+v\n", url)
	return
}