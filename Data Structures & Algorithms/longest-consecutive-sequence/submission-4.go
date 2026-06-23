func longestConsecutive(nums []int) int {
	seen := map[int]struct{}{}
	for _, n := range nums {
		seen[n] = struct{}{}
	}
	out := 0
	for _, n := range nums {
		// is this numbers -1 available?
		_, exist := seen[n-1]
		// if it doesn't exist, we're are the beginning
		// of a run
		if !exist {
			run := 1
			// keep our run alive
			_, exist = seen[n+run]
			for exist {
				run++
				_, exist = seen[n+run]
			}
			if run > out {
				out = run
			}
		}
	}
	return out
}
