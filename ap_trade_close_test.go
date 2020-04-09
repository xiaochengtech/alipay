package alipay

import (
	"fmt"
	"testing"
)

func TestTradeClose(t *testing.T) {
	fmt.Println("----------统一收单交易关闭接口----------")
	// 初始化参数
	body := TradeCloseBody{}
	body.OutTradeNo = "GYWX201908221240350122"
	// 请求支付
	aliRsp, err := testClient.TradeClose(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
