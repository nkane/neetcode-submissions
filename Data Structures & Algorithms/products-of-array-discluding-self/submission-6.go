func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	suffix := 1
	for i := n-1; i >= 0; i-- {
		prefix[i] *= suffix
		suffix *= nums[i]
	}
	return prefix
}