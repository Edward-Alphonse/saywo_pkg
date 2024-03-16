package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 将字符串生成MD5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// Md5Byte 将byte转成MD5
func Md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// Password 将一个字符串转成一个混淆密码
func Password(pwd string) string {
	md := Md5(pwd)
	suffix, prefix := md[0:8], md[24:32]
	return Md5(prefix + md + suffix)
}
