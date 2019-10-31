package alipay

import (
	"fmt"
	"testing"
)

func TestCreateTrade(t *testing.T) {
	fmt.Println("----------统一收单交易创建接口----------")
	// 初始化参数
	body := CreateTradeBody{}
	body.OutTradeNo = "GYWX201910311240354444" // TradeNo:2019082222001485841000029225
	body.SellerId = "2088102178986262"
	body.TotalAmount = 2.00
	body.Subject = "测试车场阿里支付-停车费"
	body.BuyerId = "2088102179285843"
	notifyUrl := "http://dev-service.cockoopark.com"

	// 请求支付
	aliRsp, err := testClient.CreateTrade(body, notifyUrl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
