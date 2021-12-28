package stringutil

import (
	"fmt"
	"strings"
)

func JoinAny[V any](arr []V, format, sep string) string {
	sb := strings.Builder{}
	for i, v := range arr {
		if i < 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(fmt.Sprintf(format, v))
	}
	return sb.String()
}
