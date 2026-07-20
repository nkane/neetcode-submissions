func threeSum(nums []int) [][]int {
	output := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				output = append(output, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return output
}