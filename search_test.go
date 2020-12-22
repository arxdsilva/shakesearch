package main

import (
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
