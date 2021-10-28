package tools

import (
	"strings"
	"time"
)

const gap = "#$%%$#"

// 通过加入特定界符来编码字符串数组
func EncodeStrArr(input []string) string {
	ret := ""
	max := len(input)

	ret += input[0]
	for i := 1; i < max; i++ {
		ret += gap
		ret += input[i]
	}

	return ret
}

// 将编码的字符串解码
func DecodeStrArr(input string) []string {
	ret := strings.Split(input, gap)
	return ret
}

// 将时间转化为合适的字符串
func DecodeTime(input time.Time) string {
	return input.Format("2006-01-02 15:04:05")
}
