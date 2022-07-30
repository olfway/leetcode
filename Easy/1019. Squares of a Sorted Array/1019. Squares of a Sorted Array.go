func sortedSquares(nums []int) []int {
    x, y := 0, len(nums)-1
    n := 0
    i := y
    output := make([]int, i+1)
    for i >= 0 {
        if nums[x] * nums[x] >= nums[y] * nums[y] {
            n = nums[x] * nums[x]
            x += 1
        } else {
            n = nums[y] * nums[y]
            y -= 1
        }
        output[i] = n
        i -= 1
    }
    return output
}
