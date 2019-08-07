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
	"fmt"
	"hash"
	"sort"
)

// 获取参数签名
func (c *Client) getSign(body BodyMap, signType string, privateKey string) (sign string, err error) {
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
	// 调用算法
	signStr := c.sortSignParams(body)
	_, err = h.Write([]byte(signStr))
	if err != nil {
		return
	}
	encryptedBytes, err := rsa.SignPKCS1v15(rand.Reader, key, hashs, h.Sum(nil))
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// 获取根据Key排序后的请求参数字符串
func (c *Client) sortSignParams(body BodyMap) string {
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
