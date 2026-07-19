func isPalindrome(s string) bool {
	for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
		for !isAlphaNum(s[l]) && l < r { l++ }
		for !isAlphaNum(s[r]) && l < r { r-- }
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
	}
	return true
}

func isAlphaNum(b byte) bool {
	r := rune(b)
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}