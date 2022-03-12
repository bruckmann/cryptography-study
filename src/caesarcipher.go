package src

func CeaserCipher(r rune, shift int) rune {

	// n[0] == a[0] = n[0] = a[3]

	s := int(r) + shift

	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}
