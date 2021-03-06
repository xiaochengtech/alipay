package alipay

import (
	"fmt"
	"testing"
)

func TestTradeCreate(t *testing.T) {
	fmt.Println("----------统一收单交易创建接口----------")
	// 初始化参数
	body := TradeCreateBody{}
	body.OutTradeNo = "GYWX201910311240354444"
	body.SellerId = "2088102178986262"
	body.TotalAmount = 2.00
	body.Subject = "测试车场阿里支付-停车费"
	body.BuyerId = "2088102179285843"
	notifyUrl := "http://www.example.com"
	// 请求支付
	aliRsp, err := testClient.TradeCreate(body, notifyUrl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
