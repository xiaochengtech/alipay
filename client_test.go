package alipay

import (
	"os"
)

// 测试需要设置的环境变量
var (
	AliAppId        = os.Getenv("AliAppId")
	AliPublicKey    = os.Getenv("AliPublicKey")
	AliPrivateKey   = os.Getenv("AliPrivateKey")
	AliAppAuthToken = os.Getenv("AliAppAuthToken")
)

//  测试用客户端
var testClient = NewClient(true, AliPublicKey, AliPrivateKey, Config{
	AppId:        AliAppId,
	Format:       FormatJson,
	Charset:      CharSetUTF8,
	SignType:     SignTypeRSA2,
	Version:      Version1,
	AppAuthToken: AliAppAuthToken,
})
