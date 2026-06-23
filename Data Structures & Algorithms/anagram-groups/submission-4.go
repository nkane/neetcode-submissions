func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{0}
		for _, char := range str {
			key[char-'a']++
		}
		groups[key] = append(groups[key], str)
	}
	out := make([][]string, len(groups))
	i := 0
	for _, group := range groups {
		out[i] = group
		i++
	}
	return out
}
