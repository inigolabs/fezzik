package common

import "unicode"

// UppercaseFirstChar uppercases the first char of a string
func UppercaseFirstChar(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
