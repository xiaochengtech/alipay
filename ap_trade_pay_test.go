package alipay

import (
	"fmt"
	"testing"
)

func TestAuthcode(t *testing.T) {
	code := "30763443825664394"
	ok := IsValidAuthCode(code)
	t.Logf("%s is : %v", code, ok)
}

func TestPayTrade(t *testing.T) {
	fmt.Println("----------统一收单交易支付接口----------")
	// 初始化参数
	body := PayTradeBody{}
	body.OutTradeNo = "ZSCS201910301300003333"
	body.Scene = SceneByBar
	body.AuthCode = "281778064797045892"
	body.Subject = "测试车场阿里支付-停车费"
	body.TotalAmount = 0.01
	// 请求支付
	aliRsp, err := testClient.PayTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
