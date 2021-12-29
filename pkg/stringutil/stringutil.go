package stringutil

import (
	"fmt"
	"strings"
)

type Colour string

const (
	ColourReset  Colour = "\033[0m"
	ColourRed    Colour = "\033[31m"
	ColourGreen  Colour = "\033[32m"
	ColourYellow Colour = "\033[33m"
	ColourBlue   Colour = "\033[34m"
	ColourPurple Colour = "\033[35m"
	ColourCyan   Colour = "\033[36m"
	ColourWhite  Colour = "\033[37m"
)

func Colored(s string, colour Colour) string {
	return string(colour) + s + string(ColourReset)
}

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
