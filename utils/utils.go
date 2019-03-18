package utils

import (
	"fmt"
	"strings"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
func FirstCharLower(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

func GetFirstChar(s string) string {
	return s[0:1]
}
