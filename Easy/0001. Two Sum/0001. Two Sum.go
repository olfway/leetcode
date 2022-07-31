func twoSum(nums []int, target int) []int {

    var i, n int
    length := len(nums)
    h := make(map[int]int, length)

    for i = 0; i < length; i++ {
        n = nums[i]
        if j, ok := h[target - n]; ok {
            return []int{j, i}
        }
        h[n] = i
    }
    return []int{}
}
