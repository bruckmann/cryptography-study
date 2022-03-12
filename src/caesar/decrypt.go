package caesar

// There are better ways to write this algorithm, but for educational purpose we will do this way

func Decrypt(cipherText *string) string {

	// n[0] == a[0] = n[0] = a[0 - 3]

	chars := []rune(*cipherText)
	shiftValue := 3

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var cleanText string

	for i := 0; i < len(chars); i++ {

		if string(chars[i]) == " " {
			cleanText += " "
			continue
		}

		for j := 0; j < len(alphabet); j++ {

			if string(chars[i]) == alphabet[j] {
				if j-shiftValue < 0 {
					replacementValue := len(alphabet) - (shiftValue)
					cleanText += alphabet[replacementValue]
					continue
				}
				cleanText += alphabet[j-shiftValue]
			}
		}
	}
	return cleanText
}
