func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string, 0)
	for _, s := range strs {
		var key [26]int
		for _, c := range s {
			key[c-'a']++
		}
		groups[key] = append(groups[key], s)
	}
	out := make([][]string, 0, len(groups))
	for _, v := range groups {
		out = append(out, v)
	}
	return out
}
