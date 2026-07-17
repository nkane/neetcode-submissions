func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for _, n := range nums {
        set[n] = struct{}{}
    }
    out := 0
    for _, n := range nums {
        if _, ok := set[n-1]; ok {
            continue // n-1 present → n is mid-run, skip
        }
        // n is a true start
        current := n
        run := 1
        for {
            if _, ok := set[current+1]; !ok {
                break
            }
            current++
            run++
        }
        if run > out {
            out = run
        }
    }
    return out
}
