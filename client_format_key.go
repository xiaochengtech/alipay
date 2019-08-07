package alipay

import "bytes"

// 将私钥字符串转换为RSA私钥格式
func (c *Client) FormatPrivateKey(privateKey string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen, keyLen := 64, len(privateKey)
	raws, temp := keyLen / rawLen, keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start, end := 0, 0 + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	return buffer.String()
}
