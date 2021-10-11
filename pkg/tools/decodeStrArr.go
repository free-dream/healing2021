package tools

import "strings"

const gap = "#$%%$#"

func EncodeStrArr(input []string)  string{
	var ret string

	for _, strTmp := range input {
		ret += strTmp
		ret += gap
	}
	return ret
}

func DecodeStrArr(input string)  []string{
	ret := strings.Split(input, gap)
	return ret
}
