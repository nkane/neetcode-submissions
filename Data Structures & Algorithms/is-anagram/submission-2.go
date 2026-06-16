func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var key [26]int
	for i := 0; i < len(s); i++ {
		key[s[i]-'a']++
		key[t[i]-'a']--
	}
	for _, n := range key {
		if n != 0 {
			return false
		}
	}
	return true
}