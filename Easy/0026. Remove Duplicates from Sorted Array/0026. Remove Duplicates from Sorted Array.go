func removeDuplicates(nums []int) int {
    p := 0
    length := len(nums)
    for i := 1; i < length; i++ {
        if nums[i] == nums[p] {
            continue
        }
        p++
        nums[p] = nums[i]
    }
    return p+1
}
