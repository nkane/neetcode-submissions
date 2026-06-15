func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}
	res := 0
	for _, n := range nums {
		// check if start of a sequence
		if _, ok := set[n-1]; !ok {
			run := 0
			_, exist := set[n+run]
			for exist {
				run++
				_, exist = set[n+run]
			}
			if run > res {
				res = run
			}
		}
	}
	return res
}
