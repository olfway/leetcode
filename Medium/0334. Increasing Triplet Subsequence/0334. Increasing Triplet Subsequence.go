func increasingTriplet(nums []int) bool {
    a, b := math.MaxInt32, math.MaxInt32
    for _, n := range nums {
        if n <= a { a = n; continue }
        if n <= b { b = n; continue }
        return true
    }
    return false
}

