func topKFrequent(nums []int, k int) []int {
    count := map[int]int{}
    for _, n := range nums {
        count[n]++
    }
    buckets := make([][]int, len(nums)+1)
    for n, freq := range count {
        buckets[freq] = append(buckets[freq], n)
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