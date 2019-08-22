package alipay

import (
	"fmt"
	"testing"
)

func TestRefundQueryTrade(t *testing.T) {
	fmt.Println("----------统一收单交易退款查询接口----------")
	// 初始化参数
	body := RefundQueryTradeBody{}
	body.OutTradeNo = "GYWX201908221240350111"
	body.OutRequestNo = "2019082222001485841000029220"

	// 请求支付
	aliRsp, err := testClient.RefundQueryTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
