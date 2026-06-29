func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlnum(s[i]) {
			for i < j && !isAlnum(s[i]) {
				i++
			}
		}
		if !isAlnum(s[j]) {
			for i < j && !isAlnum(s[j]) {
				j--
			}
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlnum(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}