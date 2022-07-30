func sortArrayByParity(nums []int) []int {

    length := len(nums)

    odd := 0
    even := length-1
    tmp := 0

    var found bool

    for {
        found = false
        for odd = odd; odd < length; odd++ {
            if nums[odd] % 2 != 0 {
                found = true
                break
            }
        }
        if found == false {
            return nums
        }

        found = false
        for even = even; even >= 0; even-- {
            if nums[even] % 2 == 0 {
                found = true
                break
            }
        }
        if found == false {
            return nums
        }

        if odd > even {
            return nums
        }

        tmp = nums[odd]
        nums[odd] = nums[even]
        nums[even] = tmp
    }
}
