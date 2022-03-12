package src

//There are better ways to write this algorithm, but for educational purposes, we will do this way

func CeaserCipher(cleanText *string) string {

	// n[0] == a[0] = n[0] = a[0 + 3]

	chars := []rune(*cleanText)
	shiftValue := 3

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var cipherText string

	for i := 0; i < len(chars); i++ {

		if string(chars[i]) == " " {
			cipherText += " "
			continue
		}

		for j := 0; j < len(alphabet); j++ {

			if string(chars[i]) == alphabet[j] {
				if j+shiftValue > len(alphabet) {
					replacementValue := j + shiftValue - len(alphabet)
					cipherText += alphabet[replacementValue]
					continue
				}
				cipherText += alphabet[j+shiftValue]
			}
		}
	}
	return cipherText
}
