package alipay

import (
	"fmt"
	"testing"
)

func TestTradeCancel(t *testing.T) {
	fmt.Println("----------统一收单交易撤销接口----------")
	// 初始化参数
	body := TradeCancelBody{}
	body.OutTradeNo = "GYWX201908221240350122"
	// 请求支付
	aliRsp, err := testClient.TradeCancel(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
