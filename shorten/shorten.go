package shorten

import (
	"crypto/md5"
	"encoding/binary"
	"regexp"
	"strings"
)

const base62Chars = "08ABCDEFabc1defg7hij3klmn6opqrst5uv2wxyzGHIJKLMNOP9QRSTU4VWXYZ"
const shortCodeLen = 8

var urlPathPattern = regexp.MustCompile(`^https?://[^/]+\.[^/]+/(.+)$`)

func EncodeBase62(url string) string {
	m := urlPathPattern.FindStringSubmatch(url)
	if len(m) < 2 {
		return ""
	}
	target := m[1]

	// 对路径做哈希并转成 Base62，再固定长度为 8 位
	hash := md5.Sum([]byte(target))
	n := binary.BigEndian.Uint64(hash[:8])
	code := uint64ToBase62(n)

	// 如果生成的短码长度超过8位，截取前8位；如果不足8位，前面补零
	if len(code) >= shortCodeLen {
		return code[:shortCodeLen]
	}

	return strings.Repeat(string(base62Chars[0]), shortCodeLen-len(code)) + code // 前面补零

}

func uint64ToBase62(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	buf := make([]byte, 0, 11) // 62^11 > 2^64，足够存储 uint64 转换后的结果
	for num > 0 {
		remainder := num % 62
		buf = append([]byte{base62Chars[remainder]}, buf...) // 前面插入字符
		num /= 62
	}

	return string(buf)
}
