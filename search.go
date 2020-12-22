package main

import (
	"strings"
	"unicode"
)

// spliter divides the text into words and removes ponctuation
func splitter(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
