func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}
	res := 0
	for _, n := range nums {
		// check if the current number has an available previous value
		// if not, we're are the start of a sequence
		if _, ok := set[n-1]; !ok {
			// tracks our current sequence run count total
			run := 0
			// check if the next number in sequence exist
			_, exist := set[n+run]
			for exist {
				run++
				// keep checking for next number
				_, exist = set[n+run]
			}
			if run > res {
				// set the result to the longest run
				res = run
			}
		}
	}
	return res
}
