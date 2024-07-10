package encode_and_decode_strings

import (
	"strconv"
	"strings"
)

func Encode(strs []string) string {
	serialized := make([]string, 0, len(strs)*2)
	for _, s := range strs {
		serialized = append(serialized, strconv.FormatInt(int64(len(s)), 10)+";")
		serialized = append(serialized, s)
	}

	return strings.Join(serialized, "")
}

func Decode(str string) []string {
	var strs []string

	runes := []rune(str)
	runesLen := int64(len(runes))
	for i := int64(0); i < runesLen; {
		var j int64
		for j = i; j < runesLen && runes[j] != ';' && runes[j] >= '0' && runes[j] <= '9'; j++ {
		}

		if runes[j] != ';' {
			panic("invalid format")
		}

		wordLen, err := strconv.ParseInt(string(runes[i:j]), 10, 64)
		if err != nil {
			panic(err)
		}

		i = j + 1
		strs = append(strs, string(runes[i:i+wordLen]))
		i += wordLen
	}

	return strs
}
