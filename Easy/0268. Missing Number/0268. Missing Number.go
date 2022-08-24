func missingNumber(nums []int) int {
    var n int = len(nums)
    for i, v := range nums {
        // Sum
        // n += + i - v
        // XOR
        n ^= (i ^ v)
    }
    return n
}
