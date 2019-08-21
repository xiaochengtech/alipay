package alipay

import (
	"fmt"
	"testing"
)

func TestCreateTrade(t *testing.T) {
	fmt.Println("----------统一收单交易创建接口----------")
	// 初始化参数
	body := CreateTradeBody{}
	body.OutTradeNo = "GYWX201901301040355706100456"
	body.TotalAmount = 101.00
	body.Subject = "测试车场阿里支付-停车费"
	body.BuyerId = "2088102168654131"

	// 请求支付
	aliRsp, err := testClient.CreateTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
