package util

import (
	"strings"
	"unicode"
)

func UnCapitalize(word string) string {
	firstLetter := true
	return strings.Map(func(r rune) rune {
		if firstLetter {
			firstLetter = false
			return unicode.ToLower(r)
		}
		return r
	}, word)
}
