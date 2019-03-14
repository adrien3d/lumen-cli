package utils

import (
	"strings"
	"unicode"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func CamelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

// accountName -> account name
func CamelToOriginal(s string) string {
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			words = append(words, strings.ToLower(s[lastPos:i]))
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, strings.ToLower(s[lastPos:]))
	}

	return strings.Join(words, " ")
}

func GetFirstChar(s string) string {
	return s[0:1]
}
