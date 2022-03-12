package tests

import (
	"testing"

	c "github.com/bruckmann/cryptography-study/src"
)

func TestCaesarCipher(t *testing.T) {
	cleanText := "test the caesar cipher"
	cipherText := "whvw wkh fdhvdu flskhu"

	got := c.CeaserCipher(&cleanText)

	if got != cipherText {
		t.Errorf("Want %s but received %s", cipherText, got)
	}
}
