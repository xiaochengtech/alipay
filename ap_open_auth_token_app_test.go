package alipay

import (
	"fmt"
	"testing"
)

func TestOpenAuthTokenApp(t *testing.T) {
	fmt.Println("----------换取应用授权令牌----------")
	// 初始化参数
	body := OpenAuthTokenAppBody{}
	body.GrantType = ""
	body.Code = ""
	// 请求支付
	aliRsp, err := testClient.OpenAuthTokenApp(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
