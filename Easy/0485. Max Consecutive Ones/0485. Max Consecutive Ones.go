func findMaxConsecutiveOnes(nums []int) int {
    cur := 0
    max := 0
    for _, n := range nums {
        if n == 1 {
            cur += 1
            continue
        }
        if cur > max {
            max = cur
        }
        cur = 0
    }
    if cur > max {
        max = cur
    }
    return max
}
