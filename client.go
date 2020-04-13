package alipay

type Client struct {
	config     Config // 配置信息
	publicKey  string // 支付宝公钥
	privateKey string // 应用私钥
	isProd     bool   // 是否是生产环境
}

// 初始化支付宝客户端
func NewClient(isProd bool, publicKey string, privateKey string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.publicKey = publicKey
	client.privateKey = privateKey
	client.isProd = isProd
	return client
}
