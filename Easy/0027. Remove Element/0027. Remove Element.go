func removeElement(nums []int, val int) int {
    found := 0
    length := len(nums)

    remove := 0
    keep := length - 1

    var ok bool
    var tmp int

    for {
        ok = false
        for keep = keep; keep >= 0; keep-- {
            if nums[keep] != val {
                ok = true
                keep--
                found++
                break
            }
        }
        if ok == false {
            return found
        }

        ok = false
        for remove = remove; remove <= keep; remove++ {
            if nums[remove] == val {
                ok = true
                break
            }
        }
        if ok {
            found--
            tmp = nums[keep+1]
            nums[keep+1] = nums[remove]
            nums[remove] = tmp
        }
    }
}
