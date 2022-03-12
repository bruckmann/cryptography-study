package tests

import (
	"testing"

	c "github.com/bruckmann/cryptography-study/src/caesar"
)

var cleanText string = "test the caesar cipher"
var cipherText string = "whvw wkh fdhvdu flskhu"

func TestCaesarCipherCrypt(t *testing.T) {
	got := c.Crypt(&cleanText)
	if got != cipherText {
		t.Errorf("Want %s but received %s", cipherText, got)
	}
}

func TestCaesarCipherDecrypt(t *testing.T) {
	got := c.Decrypt(&cipherText)
	if got != cleanText {
		t.Errorf("Want %s but received %s", cleanText, got)
	}
}
