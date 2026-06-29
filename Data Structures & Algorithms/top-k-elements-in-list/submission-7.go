func topKFrequent(nums []int, k int) []int {
    counts := map[int]int{}
    for _, n := range nums {
        counts[n]++
    }
    freqBucket := make([][]int, len(nums) + 1)
    for num, freq := range counts {
        freqBucket[freq] = append(freqBucket[freq], num)
    }
    output := []int{}
    for i := len(freqBucket) - 1; i >= 0; i-- {
        for _, num := range freqBucket[i] {
            output = append(output, num)
            if len(output) == k {
                return output
            }
        }
    }
    return output
}