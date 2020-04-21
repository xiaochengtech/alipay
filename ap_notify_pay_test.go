package alipay

import (
	"fmt"
	"testing"
)

func TestClientNotifyPay(t *testing.T) {
	fmt.Println("----------支付结果通知参数处理----------")
	tests := []struct {
		bodyStr string
		wantErr bool
	}{
		{"gmt_create=2020-04-21+13%3A31%3A21&charset=utf-8&seller_email=522576898%40qq.com&subject=%E4%BA%ACTEST10-%E5%81%9C%E8%BD%A6%E8%B4%B9%E7%94%A8&sign=OwsBzJSdJYDoKTUtAn2EQS8%2Bdb%2BLa97xsjx6btPjZvpLDdB%2FoyiZb2PMWIjyXWiTsaPXV2%2FzgbzW3X1vkLNqK5cZ3MzcVr8D4GhGzHGaHrwLNJBWaBqHPN0VAu5iNurOrx%2FB1No2wseDnPKW6KFTOybzLhNbzjqIUTN%2BlP6qyFixN%2FyTBSdXVbiq2%2FnK9grt8ZJKlYMt1wDgHrhoLXf3L6bJA6tnrTsDiJAY960Vmse5EzkrrpbgrlrXa71qaZOYBKwZXyxRb%2F6ACn2wcTiPnnp3wMCId4FBQm3kt4inncJYNn52lMKP%2FRSYpUWI71003uHeogJCnpjCzRHrdglO6A%3D%3D&buyer_id=2088432648838271&invoice_amount=0.01&notify_id=2020042100222133131038271411827812&fund_bill_list=%5B%7B%22amount%22%3A%220.01%22%2C%22fundChannel%22%3A%22ALIPAYACCOUNT%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.01&buyer_pay_amount=0.01&app_id=2021001155679842&sign_type=RSA2&seller_id=2088531916765168&gmt_payment=2020-04-21+13%3A31%3A31&notify_time=2020-04-21+13%3A31%3A31&version=1.0&out_trade_no=200421133120805071464420&total_amount=0.01&trade_no=2020042122001438271440382718&auth_app_id=2019080666114339&buyer_logon_id=134****5354&point_amount=0.00", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var params NotifyPayParams
			_, err := testClient.payNotifyParseParams(tt.bodyStr, &params)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyPay error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%+v\n", params)
		})
	}
}
