func thirdMax(nums []int) int {
    max1 := -1 << 63
    max2 := -1 << 63
    max3 := -1 << 63
    n := 0
    length := len(nums)
    for i := 0; i < length; i++ {
        n = nums[i]
        if n == max1 || n == max2 || n == max3 {
            continue
        }
        if n > max1 {
            max3 = max2
            max2 = max1
            max1 = n
            continue
        }
        if n > max2 {
            max3 = max2
            max2 = n
            continue
        }
        if n > max3 {
            max3 = n
            continue
        }
    }
    if max3 > -1 << 63 {
        return int(max3)
    }
    return int(max1)
}
