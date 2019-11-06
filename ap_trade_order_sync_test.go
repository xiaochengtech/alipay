package alipay

import (
	"fmt"
	"testing"
)

func TestOrderSyncTrade(t *testing.T) {
	fmt.Println("----------支付宝订单信息同步接口----------")
	// 初始化参数
	body := OrderSyncTradeBody{}
	body.TradeNo = "GYWX201908221240350111"
	body.OutRequestNo = "GYWX201908221240350111"
	body.BizType = "CREDIT_AUTH"
	// 请求支付
	aliRsp, err := testClient.OrderSyncTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
