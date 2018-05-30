package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

func MakeSignature(token, timestamp, nonce string) string {
	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}
