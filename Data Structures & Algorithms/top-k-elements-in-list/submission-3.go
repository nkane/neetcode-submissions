func topKFrequent(nums []int, k int) []int {
    counts := map[int]int{}
    for _, n := range nums {
        counts[n]++
    }
    buckets := make([][]int, len(nums)+1)
    for value, freq := range counts {
        buckets[freq] = append(buckets[freq], value)
    }
    out := make([]int, k)
    j := 0
    for i := len(buckets)-1; i >= 0; i-- {
        for _, value := range buckets[i] {
            if j == k {
                return out
            }
            out[j] = value
            j++
        }
    }
    return out
}
