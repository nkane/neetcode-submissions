func topKFrequent(nums []int, k int) []int {
    counts := map[int]int{}
    for _, num := range nums {
        counts[num]++
    }
    freq := make(map[int][]int, 0)
    for num, count := range counts {
        freq[count] = append(freq[count], num)
    }
    buckets := make([][]int, len(nums)+1)
    for k, v := range freq {
        buckets[k] = v
    }
    out := make([]int, 0)
    idx := len(buckets) - 1
    n := len(out)
    for n < k {
        for _, item := range buckets[idx]  {
            out = append(out, item)
            if len(out) == k {
                return out
            }
        }
        n = len(out)
        idx--
    }
    return out
}