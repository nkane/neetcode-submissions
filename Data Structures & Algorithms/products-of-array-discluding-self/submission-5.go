func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n-2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	output := make([]int, n)
	for i := 0; i < n; i++ {
		output[i] = prefix[i] * suffix[i]
	}
	return output
}