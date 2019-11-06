package alipay

import (
	"fmt"
	"testing"
)

func TestRefundTrade(t *testing.T) {
	fmt.Println("----------统一收单交易退款接口----------")
	// 初始化参数
	body := RefundTradeBody{}
	body.OutTradeNo = "ZSCS201910301300003333"
	body.RefundAmount = 0.01
	// 请求支付
	aliRsp, err := testClient.RefundTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
