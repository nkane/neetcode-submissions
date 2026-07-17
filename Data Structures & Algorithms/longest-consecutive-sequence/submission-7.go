func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}
	out := 0
	for _, n := range nums {
		_, exist := set[n-1]
		run := 1
		for exist {
			_, exist = set[n+run]
			run++
		}
		if out < run {
			out = run
		}
	}
	return out
}