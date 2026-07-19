func twoSum(nums []int, target int) []int {
	out := []int{}
	for l, r := 0, len(nums)-1; l < r; {
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			out = append(out, l+1)
			out = append(out, r+1)
			break
		}
	}
	return out
}
