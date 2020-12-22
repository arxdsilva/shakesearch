package main

import (
	"strings"
	"unicode"

	snowball "github.com/kljensen/snowball/english"
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

// 100 common words in eng
// we'll drop it to prevent getting too much unwanted
// https://www.espressoenglish.net/the-100-most-common-words-in-english/
var commonWords = map[string]string{
	"the": "", "at": "", "there": "", "some": "", "my": "",
	"of": "", "be": "", "use": "", "her": "", "than": "",
	"and": "", "this": "", "an": "", "would": "", "first": "",
	"a": "", "have": "", "each": "", "make": "", "water": "",
	"to": "", "from": "", "which": "", "like": "", "been": "",
	"in": "", "or": "", "she": "", "him": "", "call": "",
	"is": "", "one": "", "do": "", "into": "", "who": "",
	"you": "", "had": "", "how": "", "time": "", "oil": "",
	"that": "", "by": "", "their": "", "has": "", "its": "",
	"it": "", "word": "", "if": "", "look": "", "now": "",
	"he": "", "but": "", "will": "", "two": "", "find": "",
	"was": "", "not": "", "up": "", "more": "", "long": "",
	"for": "", "what": "", "other": "", "write": "", "down": "",
	"on": "", "all": "", "about": "", "go": "", "day": "",
	"are": "", "were": "", "out": "", "see": "", "did": "",
	"as": "", "we": "", "many": "", "number": "", "get": "",
	"with": "", "when": "", "then": "", "no": "", "come": "",
	"his": "", "your": "", "them": "", "way": "", "made": "",
	"they": "", "can": "", "these": "", "could": "", "may": "",
	"I": "", "said": "", "so": "", "people": "", "part": ""}

// removeCommonWords helps the search to be more meaningful
// avoiding common words used in eng
func removeCommonWords(txtSet []string) (rdcw []string) {
	rdcw = make([]string, 0, len(txtSet))
	for _, word := range txtSet {
		if _, ok := commonWords[word]; !ok {
			rdcw = append(rdcw, word)
		}
	}
	return
}

// stemWords reduces the words to its stem base form,
// eg: auditor > audit
func stemWords(txtSet []string) (stm []string) {
	stm = make([]string, len(txtSet))
	for pos, txt := range txtSet {
		stm[pos] = snowball.Stem(txt, false)
	}
	return
}

func filterText(text string) (filtered []string) {
	filtered = splitter(text)
	filtered = toLower(filtered)
	filtered = removeCommonWords(filtered)
	filtered = stemWords(filtered)
	return
}
