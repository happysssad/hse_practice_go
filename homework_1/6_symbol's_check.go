func isVowel(letter byte) bool {
	vowels := [10]byte{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	for _, b := range vowels {
		if letter == b {
			return true
		}
	}
	return false
}
