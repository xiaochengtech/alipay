package alipay

import (
	"fmt"
	"testing"
)

func TestOpenAuthTokenApp(t *testing.T) {
	fmt.Println("----------换取应用授权令牌----------")
	// 初始化参数
	body := OpenAuthTokenAppBody{
		GrantType: GrantTypeAuthorizationCode,
		Code:      "78241ee2c80848a3aec1f82c79eb3X98",
	}
	// 请求支付
	aliRsp, err := testClient.OpenAuthTokenApp(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", aliRsp)
}
