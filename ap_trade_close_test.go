package alipay

import (
	"fmt"
	"testing"
)

func TestCloseTrade(t *testing.T) {
	fmt.Println("----------统一收单交易关闭接口----------")
	// 初始化参数
	body := CloseTradeBody{}
	body.OutTradeNo = "GYWX201908221240350111"
	//body.TradeNo = "2019082222001454131000044049"

	// 请求支付
	aliRsp, err := testClient.CloseTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
