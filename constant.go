package alipay

type BodyMap map[string]interface{}

const (
	baseUrl        = "https://openapi.alipay.com/gateway.do"    // (生产环境) 支付宝接口地址
	baseUrlSandbox = "https://openapi.alipaydev.com/gateway.do" // (沙盒环境) 支付宝接口地址

	// 测试用公钥密钥
	aliPayPublicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzd66NCUBXvqurS1KlWlcVtu6oZoYk5jwlWB1E2h+3XaXYZxAOrr0NRTmL83jN5NwxwT2btf3LNhl0UOWLiXk7wYG2rYUYhGDp4LF7V0pkkZJ0DvrfYTWkLHYw9w5NJ9sspTme8I+Rs6000J/89Ar5ms6IFIywP4dJyTTEMeByDqIh7qyPckU4yI6kuqshamP21Ssr/GgJ0VLgSO9GlCEEhTwth8mWSH5bhuUSobWEMDbIbC3r+KL3nHSJQ0h/9lVvJBNK0erFJ8mxJbxNpqXlXozXfbpScikT7d4b4mcAXPJFlXifUUfyOzwpRW6IC4l6VGeMj7nYtwwECc2a6gS3wIDAQAB"
	aliPayPrivateKey = "MIIEogIBAAKCAQEAzd66NCUBXvqurS1KlWlcVtu6oZoYk5jwlWB1E2h+3XaXYZxAOrr0NRTmL83jN5NwxwT2btf3LNhl0UOWLiXk7wYG2rYUYhGDp4LF7V0pkkZJ0DvrfYTWkLHYw9w5NJ9sspTme8I+Rs6000J/89Ar5ms6IFIywP4dJyTTEMeByDqIh7qyPckU4yI6kuqshamP21Ssr/GgJ0VLgSO9GlCEEhTwth8mWSH5bhuUSobWEMDbIbC3r+KL3nHSJQ0h/9lVvJBNK0erFJ8mxJbxNpqXlXozXfbpScikT7d4b4mcAXPJFlXifUUfyOzwpRW6IC4l6VGeMj7nYtwwECc2a6gS3wIDAQABAoIBAH3ySvxnv0T6HC86TCzIPoOs/aqr+Wki4dyhHD75aNPYH5jJ3MmRYqGu3XxVezKe8xnuwUzitwcqhG/X52LoL/zDNSJMzn+Is4ChkrE6k5o3giTw44rw38u54iFA84ZsGjXOvorsNOlgwGOWhW0F6DWZWAA6CW7kh7VQ5AeZ5p41iAXyHcEhgXtg1xVyC4hahownrhJvUB2sygpt23yPXLNq0so/2TucZkes1Nf43pN9PfJ+eIYLupryfxYm454DISR7jtPMgA/tF28mjGc5fheV/+r9Q+sLRJCWxVqYt0SqsLJxV7HwvA7HkChwy7WALyXvyrBZbOP99IY0RQJq+HECgYEA7kwSObQbTI4zFQMZ3WzUiTXyXtZ9XUuWDor/IYhoF+ZbcexkskQf9WsbZQmj4vA0sGRrqpCp5BC7BMPpEkwQFI7FLVekx7l1wjk1D5ZXCrF2sGpjpFaUgfaxtcsdz166yNwdaoAjLV5wCIiyv74t4Q/jnqav3Fel766d7LWdHAUCgYEA3Sn1PvZ1VdSUuC9bUZ3UylQ6sKz3qS2CQ5gXuSsKSxQqwEC9ldy1E8GzXqnOfRrG0resfpR8jKMyG+xsXdLGLrQDrJtwmcxY9w2SkDtHifjkXmJ5nxQhP/ZXGwO/peH1vTDFQcjmjcDUB8SF4+femmUM/k45SOSALHVz0Fc5zJMCgYAMIUMsPMbG9tYS072VT9zhvyU/PpvTC+3vf5PNGSxzSv6MpgPaLy2RSIp0cjEtdBy6feIdJJABU/ixHWfXzpdi9IGE9Vkl8YLVCCRzTqvPl4j7Ie9Ahke49d3F6zjxPVk0sD+vFCa2QIBjbtHqgLvFIRHtGX8KMEv6D2FqkzQUIQKBgGWoX54tUwm8yQ9QVT1Suvi1drS3DK+qx3Gie2UVr4mLH4t3Nq2n4WPPaU9d5hgDMtiwrI0SW52ZdfL33WfS9l2JSRcR69QF5e7JQlEhrmRxDyeDRDwm2JG5/ZfhRTpOftITlNgELFWECPH5F0IG1n9Ja4uJETO8NF47LpzZOywZAoGALYlQ5ICfbg6wIV58Q0BTMMebUlSG94dabRMD7d6qnPcHU5JOt0i+2QsawSa3LGQRcahjdn6QcJje00DlSE3OkKv4xvmUidQ4WUV1bQJq7SHEG3aqvOWMm1caViXQ17GJygDq2pzkn8G87jqc90DMYjWGOqTOuTbGfmTh07Ac8Rg="

	// 请求格式
	FormatJson = "JSON" // Json格式

	// 编码格式
	CharSetUTF8 = "UTF-8" // UTF8

	// 签名算法类型
	SignTypeRSA  = "RSA"  // RSA签名
	SignTypeRSA2 = "RSA2" // RSA2签名

	// 版本号
	Version1 = "1.0" // 1.0版本

	// 货币类型
	FeeTypeCNY = "CNY" // 人民币

	TransInTypeCard  = "cardAliasNo" // 结算收款方的银行卡编号
	TransInTypeUser  = "userId"      // 表示是支付宝账号对应的支付宝唯一用户号
	TransInTypeLogin = "loginName"   // 表示是支付宝登录号

	LogisticsTypePost    = "POST"    // 平邮
	LogisticsTypeExpress = "EXPRESS" // 其他快递
	LogisticsTypeVirtual = "VIRTUAL" // 虚拟物品
	LogisticsTypeEms     = "EMS"     // EMS
	LogisticsTypeDirect  = "DIRECT"  // 无需物流

	SceneByBar  = "bar_code"  // 条码支付
	SceneByWave = "wave_code" // 声波支付

	AuthConfirmModeComplete    = "COMPLETE"     // 转交易支付完成结束预授权，解冻剩余金额
	AuthConfirmModeNotComplete = "NOT_COMPLETE" // 转交易支付完成不结束预授权，不解冻剩余金额
)
