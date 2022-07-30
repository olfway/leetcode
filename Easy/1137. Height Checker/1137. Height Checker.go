func heightChecker(heights []int) int {
    n := 0
    length := len(heights)
    heights2 := make([]int, length)
    copy(heights2, heights)
    sort.Ints(heights)
    for i := 0; i < length; i++ {
        if heights[i] != heights2[i] {
            n++
        }
    }
    return n
}
