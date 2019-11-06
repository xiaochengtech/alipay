package alipay

import (
	"fmt"
	"testing"
)

func TestPreCreateTrade(t *testing.T) {
	fmt.Println("----------统一收单交易预创接口----------")
	// 初始化参数
	body := PreCreateTradeBody{}
	body.OutTradeNo = "ZSCS201908221300003333"
	body.TotalAmount = 0.01
	body.Subject = "测试车场阿里支付-停车费"
	// 请求支付
	aliRsp, err := testClient.PreCreateTrade(body, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
