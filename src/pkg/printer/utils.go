package printer

import (
	"errors"
	"math"
	"strings"
)

const PADDING = 2

func maxStringLength(list []string) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("cannot process empty list")
	}

	var maxLen = -1

	for _, str := range list {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	return maxLen, nil
}

func nChar(n int, char rune) string {
	var str string

	for i := 0; i < n; i++ {
		str += string(char)
	}

	return str
}

func centerString(text string, width int) string {
	padding := width - len(text)

	if padding%2 == 0 {
		return strings.Join([]string{nChar(padding/2, ' '), text, nChar(padding/2, ' ')}, " ")
	} else {
		left := int(math.Floor(float64(padding / 2)))
		right := int(math.Floor(float64(padding/2)) + 1)
		return strings.Join([]string{nChar(left, ' '), text, nChar(right, ' ')}, " ")
	}
}
