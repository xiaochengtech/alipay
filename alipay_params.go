//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/13 14:42
//==================================
package alipay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"hash"
	"net/url"
	"sort"
)

//	AppId      string `json:"app_id"`      //支付宝分配给开发者的应用ID
//	Method     string `json:"method"`      //接口名称
//	Format     string `json:"format"`      //仅支持 JSON
//	ReturnUrl  string `json:"return_url"`  //HTTP/HTTPS开头字符串
//	Charset    string `json:"charset"`     //请求使用的编码格式，如UTF-8,GBK,GB2312等，推荐使用 UTF-8
//	SignType   string `json:"sign_type"`   //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
//	Sign       string `json:"sign"`        //商户请求参数的签名串
//	Timestamp  string `json:"timestamp"`   //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
//	Version    string `json:"version"`     //调用的接口版本，固定为：1.0
//	NotifyUrl  string `json:"notify_url"`  //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
//	BizContent string `json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档

//设置支付后的ReturnUrl
func (this *aliPayClient) SetReturnUrl(url string) (client *aliPayClient) {
	this.ReturnUrl = url
	return this
}

//设置支付宝服务器主动通知商户服务器里指定的页面http/https路径。
func (this *aliPayClient) SetNotifyUrl(url string) (client *aliPayClient) {
	this.NotifyUrl = url
	return this
}

//设置编码格式，如UTF-8,GBK,GB2312等，推荐使用 UTF-8
func (this *aliPayClient) SetCharset(charset string) (client *aliPayClient) {
	if charset == null {
		this.Charset = "UTF-8"
	} else {
		this.Charset = charset
	}
	return this
}

//设置签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
func (this *aliPayClient) SetSignType(signType string) (client *aliPayClient) {
	if signType == null {
		this.SignType = "RSA2"
	} else {
		this.SignType = signType
	}
	return this
}

//设置应用授权
func (this *aliPayClient) SetAppAuthToken(appAuthToken string) (client *aliPayClient) {
	this.AppAuthToken = appAuthToken
	return this
}
