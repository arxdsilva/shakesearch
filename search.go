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

func toLower(txtSet []string) (lc []string) {
	lc = make([]string, len(txtSet))
	for pos, txt := range txtSet {
		lc[pos] = strings.ToLower(txt)
	}
	return lc
}
