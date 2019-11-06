package alipay

import (
	"fmt"
	"testing"
)

func TestSystemOAuthToken(t *testing.T) {
	fmt.Println("----------换取授权访问令牌----------")
	// 初始化参数
	body := SystemOAuthTokenBody{}
	body.GrantType = "authorization_code"
	body.Code = "GYWX201908221240350122"
	// 请求支付
	aliRsp, err := testClient.SystemOAuthToken(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
