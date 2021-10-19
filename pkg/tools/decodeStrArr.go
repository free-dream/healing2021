package tools

import "strings"

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
