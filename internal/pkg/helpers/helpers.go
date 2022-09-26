package helpers

import "strings"

func EscapeBrackets(str string) string {
	s := strings.Replace(str, "(", "\\(", -2)
	s = strings.Replace(s, ")", "\\)", -2)
	return s
}
