func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}
	out := 0
	for _, n := range nums {
		if _, ok := set[n-1]; !ok {
			run := 0
			_, exist := set[n+run]
			for exist {
				run++
				_, exist = set[n+run]
			}
			if run > out {
				out = run
			}
		}
	}
	return out
}