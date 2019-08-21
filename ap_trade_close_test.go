package alipay

import (
	"fmt"
	"testing"
)

func TestCloseTrade(t *testing.T) {
	fmt.Println("----------统一收单交易关闭接口----------")
	// 初始化参数
	body := CloseTradeBody{}
	body.OutTradeNo = "GYWX201908211040355706100456"

	// 请求支付
	aliRsp, err := testClient.CloseTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
