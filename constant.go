package alipay

type BodyMap map[string]interface{}

const (
	appIdSandbox = "2016101000655815"

	baseUrl        = "https://openapi.alipay.com/gateway.do"    // (生产环境) 支付宝接口地址
	baseUrlSandbox = "https://openapi.alipaydev.com/gateway.do" // (沙盒环境) 支付宝接口地址

	// 请求格式
	FormatJson = "JSON" // Json格式

	// 编码格式
	CharSetUTF8 = "UTF-8" // UTF8

	// 签名算法类型
	SignTypeRSA  = "RSA"  // RSA签名
	SignTypeRSA2 = "RSA2" // RSA2签名

	// 版本号
	Version1 = "1.0" // 1.0版本

	// 授权类型
	GrantTypeAuthorizationCode = "authorization_code" // 用code获取
	GrantTypeRefreshToken      = "refresh_token"      // 用refresh_token获取

	// 货币类型
	FeeTypeCNY = "CNY" // 人民币

	// 回调通知类型
	NotifyTypeTradeStatusSync = "trade_status_sync" // 交易状态通知

	// 回调中的交易状态
	TradeStatusSuccess = "TRADE_SUCCESS" // 交易成功

	TransInTypeCard  = "cardAliasNo" // 结算收款方的银行卡编号
	TransInTypeUser  = "userId"      // 表示是支付宝账号对应的支付宝唯一用户号
	TransInTypeLogin = "loginName"   // 表示是支付宝登录号

	TransOutTypeUser  = "userId"    // 表示是支付宝账号对应的支付宝唯一用户号
	TransOutTypeLogin = "loginName" // 表示是支付宝登录号

	LogisticsTypePost    = "POST"    // 平邮
	LogisticsTypeExpress = "EXPRESS" // 其他快递
	LogisticsTypeVirtual = "VIRTUAL" // 虚拟物品
	LogisticsTypeEms     = "EMS"     // EMS
	LogisticsTypeDirect  = "DIRECT"  // 无需物流

	SceneByBar  = "bar_code"  // 条码支付
	SceneByWave = "wave_code" // 声波支付

	AuthConfirmModeComplete    = "COMPLETE"     // 转交易支付完成结束预授权，解冻剩余金额
	AuthConfirmModeNotComplete = "NOT_COMPLETE" // 转交易支付完成不结束预授权，不解冻剩余金额

	RoyaltyTypeTransfer  = "transfer"  // 普通分账
	RoyaltyTypeReplenish = "replenish" // 补差

	BizTypeCreditAuth   = "CREDIT_AUTH"   // 信用授权场景下传
	BizTypeCreditDeduct = "CREDIT_DEDUCT" // 信用代扣场景下传

	OrderBizStatusComplete = "COMPLETE" // 同步用户已履约
	OrderBizStatusClosed   = "CLOSED"   // 同步履约已取消
	OrderBizStatusViolated = "VIOLATED" // 用户已违约

	ResponseCodeSuccess             = "10000" // 调用成功
	ResponseCodeServiceNotAvaliable = "20000" // 服务不可用
	ResponseCodeNotAuthrise         = "20001" // 授权权限不足
	ResponseCodeLessParameters      = "40001" // 缺少必选参数
	ResponseCodeInvalidParameters   = "40002" // 非法的参数
	ResponseCodeFailure             = "40004" // 业务处理失败
	ResponseCodePermissionDeny      = "40006" // 权限不足
)
