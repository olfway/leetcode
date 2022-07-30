func moveZeroes(nums []int)  {
    var zeros int
    last := len(nums) - 1
    for i := 0; i < last; i++ {
        if nums[i] != 0 {
            continue
        }
        zeros = 0
        for i + zeros < last && nums[i + zeros] == 0 {
            zeros++
		}
		copy(nums[i:last - zeros + 1], nums[i + zeros:last + 1])
        last -= zeros
    }
    for last = last + 1; last < len(nums); last++ {
		nums[last] = 0
    }
}
