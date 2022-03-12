package src

func CeaserCipher(cleanText *string) string {

	// n[0] == a[0] = n[0] = a[3]
	chars := []rune(*cleanText)

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	REPLACEMENT_NUMBER := 3

	var cipherText string

	for i := 0; i < len(chars); i++ {

		if string(chars[i]) == " " {
			cipherText += " "
			continue
		}

		for j := 0; j < len(alphabet); j++ {

			if string(chars[i]) == alphabet[j] {
				if j+REPLACEMENT_NUMBER > len(alphabet) {
					replacementValue := j + REPLACEMENT_NUMBER - len(alphabet)
					cipherText += alphabet[replacementValue]
					continue
				}
				cipherText += alphabet[j+REPLACEMENT_NUMBER]
			}
		}
	}
	return cipherText
}
