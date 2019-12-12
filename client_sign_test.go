package alipay

import (
	"fmt"
	"testing"
)

func TestVerifySign(t *testing.T) {
	fmt.Println("----------异步回调的验签接口----------")
	// 初始化参数
	url := ""
	// 请求支付
	if err := testClient.VerifySign(url); err != nil {
		t.Error(err)
	} else {
		t.Log("验签成功")
	}
}
