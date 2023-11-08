func productExceptSelf(nums []int) []int {

    length := len(nums)
    out := make([]int, length)

    left := 1
    for i := 0; i < length; i++ {
        out[i] = left
        left *= nums[i]
    }

    right := 1
    for i := length-1; i >= 0; i-- {
        out[i] *= right
        right *= nums[i]
    }

    return out
}


