func findDisappearedNumbers(nums []int) []int {
	var i, n int
	length := len(nums)
	for i < length {
        n = nums[i]
        if nums[n-1] != n {
			nums[i] = nums[n-1]
			nums[n-1] = n
			continue
		}
		i++
	}
	output := make([]int, 0, length)
    for i = 0; i < length; i++ {
        if nums[i] != i+1 {
            output = append(output, i+1)
        }
    }
    return output
}
