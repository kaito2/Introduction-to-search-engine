package test

import (
	"github.com/kaito2/Introduction-to-search-engine/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testIgnoredChars = []rune{
	' ', '\f', '\n', '\r', '\t', '\v', '!', '"', '#',
	'$', '%', '&', '\'', '(', ')', '*', '+', ',', '-',
	'.', '/', ',', ';', '<', '=', '>', '?', '@', '[',
	'\\', ']', '^', '_', '`', '{', '|', '}', '~',
	'　', '、', '。', '）', '（',
}

func TestIsIgnoredChar(t *testing.T) {
	for _, c := range testIgnoredChars {
		assert.Equal(t, modules.IsIgnoredChar(c), true)
	}
}
