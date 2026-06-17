func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isAlpha(s[i]) {
			i++
		}
		for i < j && !isAlpha(s[j]) {
			j--
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlpha(c byte) bool {
	r := rune(c)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}