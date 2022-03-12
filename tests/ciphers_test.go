package tests

import (
	"testing"

	c "github.com/bruckmann/cryptography-study/src"
)

func TestCaesarCipher(t *testing.T) {
	var cleanText string = "meet me after the toga party"

	var cipherText string = "phhw ph diwhu wkh wrjd sduwb"
	got := c.CeaserCipher(&cleanText)

	if got != cipherText {
		t.Errorf("Want %s but received %s", cipherText, got)
	}
}
