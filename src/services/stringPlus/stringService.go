package stringPlus

import (
	"strings"
)

func Strtr(str string, m map[string]string) string {
	for k, v := range m {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}
