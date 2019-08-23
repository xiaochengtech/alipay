package alipay

import (
	"fmt"
	"testing"
)

func TestQueryTrade(t *testing.T) {
	fmt.Println("----------统一收单线下交易查询接口----------")
	// 初始化参数
	body := QueryTradeBody{}
	body.OutTradeNo = "ZSCS201908221300003333"

	// 请求支付
	aliRsp, err := testClient.QueryTrade(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", aliRsp)
}
