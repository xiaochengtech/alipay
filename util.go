package alipay

import "regexp"

// 25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准
func IsValidAuthCode(authcode string) (ok bool) {
	pattern := "^(2[5-9]|30)[0-9]{14,22}$"
	ok, _ = regexp.MatchString(pattern, authcode)
	return
}
