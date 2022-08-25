func intersect(nums1 []int, nums2 []int) []int {

    if len(nums2) < len(nums1) {
        nums1, nums2 = nums2, nums1
    }

    output := make([]int, 0, len(nums1))
    cache := make(map[int]int, len(nums1))

    for _, n := range nums1 {
        cache[n]++
    }

    for _, n := range nums2 {
        if cache[n] > 0 {
            cache[n]--
            output = append(output, n)
        }
    }

    return output
}
