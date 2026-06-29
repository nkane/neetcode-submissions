func longestConsecutive(nums []int) int {
	seen := map[int]struct{}{}
	for _, n := range nums {
		seen[n] = struct{}{}
	}
	count := 0
	for _, num := range nums {
		_, exist := seen[num-1]
		if !exist {
			run := 1
			_, exist = seen[num+run]
			for exist {
				run++
				_, exist = seen[num+run]
			}
			if run > count {
				count = run
			}
		}
	}
	return count
}