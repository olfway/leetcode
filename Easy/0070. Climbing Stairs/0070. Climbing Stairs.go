func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    p2, p1 := 1, 2
    for i := 3; i < n; i++ {
        p2, p1 = p1, p2 + p1
    }
    return p2 + p1
}
