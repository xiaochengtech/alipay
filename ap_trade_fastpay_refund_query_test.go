package alipay

import (
	"fmt"
	"testing"
)

func TestTradeFastpayRefundQuery(t *testing.T) {
	fmt.Println("----------统一收单交易退款查询接口----------")
	// 初始化参数
	body := TradeFastpayRefundQueryBody{}
	body.OutTradeNo = "GYWX201908221140351111"
	body.OutRequestNo = "2019082222001485841000025596"
	// 请求支付
	aliRsp, err := testClient.TradeFastpayRefundQuery(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
