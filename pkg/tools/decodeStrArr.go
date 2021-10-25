package tools

import (
	"strings"
	"time"
)

const gap = "#$%%$#"

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

func DecodeStrArr(input string) []string {
	ret := strings.Split(input, gap)
	return ret
}

func DecodeTime(input time.Time)  string{
	return input.Format("2006-01-02 15:04:05")
}
