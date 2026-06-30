func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for num, freq := range count {
		bucket[freq] = append(bucket[freq], num)
	}
	out := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, num := range bucket[i] {
			out = append(out, num)
			if len(out) == k {
				return out
			}
		}
	}
	return out
}