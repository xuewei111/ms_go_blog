package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// 给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @params str 返回的md5
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
