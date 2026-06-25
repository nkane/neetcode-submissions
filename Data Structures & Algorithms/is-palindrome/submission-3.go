func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for !isAlnum(rune(s[i])) && i < j {
			i++
		}
		for !isAlnum(rune(s[j])) && i < j {
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

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}