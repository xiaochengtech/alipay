package alipay

import (
	"fmt"
	"testing"
)

func TestCancelTrade(t *testing.T) {
	fmt.Println("----------统一收单交易撤销接口----------")
	// 初始化参数
	body := CancelTradeBody{}
	body.OutTradeNo = "GYWX201901301040355706100456"

	// 请求支付
	aliRsp, err := testClient.CancelTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
