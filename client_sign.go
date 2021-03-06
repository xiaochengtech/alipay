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
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"sort"
)

// 获取参数签名
func (c Client) getSign(body BodyMap, signType string, privateKey string) (sign string, err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
	)
	// 解析秘钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = errors.New("秘钥错误")
		return
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	// 获取加密算法
	switch signType {
	case SignTypeRSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case SignTypeRSA2:
		fallthrough
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	// 拼接原始串
	signStr := c.sortSignParams(body)
	_, err = h.Write([]byte(signStr))
	if err != nil {
		return
	}
	// 调用算法
	encryptedBytes, err := rsa.SignPKCS1v15(rand.Reader, key, hashs, h.Sum(nil))
	if err != nil {
		return
	}
	// base64转码
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// 验证同步返回值签名
func (c Client) verifySignSync(data interface{}, sign string) (err error) {
	var (
		h    hash.Hash
		hash crypto.Hash
	)
	// 签名转码
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	// 获取待签名原始串
	dataBytes, _ := json.Marshal(data)
	fmt.Println(string(dataBytes))
	// 解析公钥
	pKey := c.FormatPublicKey(c.publicKey)
	block, _ := pem.Decode([]byte(pKey))
	if block == nil {
		err = errors.New("支付宝公钥错误")
		return
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		err = errors.New("支付宝公钥转换错误")
		return
	}
	// 判断签名方式
	switch c.config.SignType {
	case SignTypeRSA:
		h = sha1.New()
		hash = crypto.SHA1
	case SignTypeRSA2:
		fallthrough
	default:
		h = sha256.New()
		hash = crypto.SHA256
	}
	// 调用签名校验
	_, _ = h.Write(dataBytes)
	hashed := h.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, hash, hashed, signBytes)
	return
}

// 验证异步返回值签名
func (c Client) verifySignAyn(data interface{}, sign string) (err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
		body  BodyMap
	)
	// 解析参数
	dataBytes, _ := json.Marshal(data)
	if err = json.Unmarshal(dataBytes, &body); err != nil {
		return
	}
	signData := c.sortSignParams(body)
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	// 解析公钥
	pKey := c.FormatPublicKey(c.publicKey)
	block, _ := pem.Decode([]byte(pKey))
	if block == nil {
		err = errors.New("支付宝公钥错误")
		return
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		err = errors.New("支付宝公钥转换错误")
		return
	}
	// 判断签名方式
	switch c.config.SignType {
	case SignTypeRSA:
		hashs = crypto.SHA1
	case SignTypeRSA2:
		fallthrough
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	_, _ = h.Write([]byte(signData))
	err = rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
	return
}

// 获取根据Key排序后的请求参数字符串
func (c Client) sortSignParams(body BodyMap) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, body[k]))
	}
	s, i := buffer.String(), buffer.Len()
	return s[:i-1]
}
