package tests

import (
	"strings"
	"testing"

	c "github.com/bruckmann/cryptography-study/src"
)

func TestCaesarCipher(t *testing.T) {
	var cleanText string = "meet me after the toga party"

	var cipherText string = "phhw=ph=diwhu=wkh=wrjd=sduwb"

	got := strings.Map(func(r rune) rune {
		return c.CeaserCipher(r, 3)
	}, cleanText)

	if got != cipherText {
		t.Errorf("Want %s but received %s", cipherText, got)
	}
}
