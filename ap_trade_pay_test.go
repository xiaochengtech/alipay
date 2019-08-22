package alipay

import (
	"fmt"
	"testing"
)

func TestPayTrade(t *testing.T) {
	fmt.Println("----------统一收单交易支付接口----------")
	// 初始化参数
	body := PayTradeBody{}
	body.OutTradeNo = "GYWX201908221240354444"
	body.Scene = SceneByBar
	body.AuthCode = "282656867596265976"
	body.Subject = "测试车场阿里支付-停车费"
	body.TotalAmount = 2.00

	// 请求支付
	aliRsp, err := testClient.PayTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
