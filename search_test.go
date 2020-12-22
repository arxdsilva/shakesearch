package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spliter(t *testing.T) {
	txt := "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab"
	splitTxt := splitter(txt)
	assert.Equal(t, 40, len(splitTxt))
	for _, tx := range splitTxt {
		assert.NotContains(t, tx, ".")
		assert.NotContains(t, tx, ",")
	}
}

func Test_toLower(t *testing.T) {
	txt := "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab"
	splitTxt := splitter(txt)
	assert.Equal(t, 40, len(splitTxt))
	for _, tx := range splitTxt {
		assert.NotContains(t, tx, ".")
		assert.NotContains(t, tx, ",")
	}
	lower := toLower(splitTxt)
	assert.Equal(t, "the", lower[0])
}

func Test_removeCommonWords(t *testing.T) {
	txt := "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab"
	splitTxt := splitter(txt)
	assert.Equal(t, 40, len(splitTxt))
	for _, tx := range splitTxt {
		assert.NotContains(t, tx, ".")
		assert.NotContains(t, tx, ",")
	}
	lower := toLower(splitTxt)
	assert.Equal(t, "the", lower[0])
	filtered := removeCommonWords(lower)
	assert.Equal(t, 34, len(filtered))
}

func Test_stemWords(t *testing.T) {
	txt := "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab"
	splitTxt := splitter(txt)
	assert.Equal(t, 40, len(splitTxt))
	for _, tx := range splitTxt {
		assert.NotContains(t, tx, ".")
		assert.NotContains(t, tx, ",")
	}
	lower := toLower(splitTxt)
	fmt.Println(lower, len(lower))
	filtered := removeCommonWords(lower)
	assert.Equal(t, 34, len(filtered))
	stemmed := stemWords(filtered)
	assert.Equal(t, "jump", stemmed[3])
	assert.Equal(t, "grace", stemmed[16])
	assert.Equal(t, "whelp", stemmed[18])
}

func Test_filterText(t *testing.T) {
	txt := "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab"
	f := filterText(txt)
	assert.Equal(t, "jump", f[3])
	assert.Equal(t, "grace", f[16])
	assert.Equal(t, "whelp", f[18])
}

func Test_removeDuplicates(t *testing.T) {
	idxs := []int{100, 200, 300, 400, 500, 600, 700}
	idxs = removeDuplicates(idxs)
	expected := []int{100, 400, 700}
	assert.Equal(t, expected, idxs)
}
