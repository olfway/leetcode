func findNumbers(nums []int) int {
    output := 0
    for _, n := range nums {
        digits := 1
        for n >= 10 {
            digits += 1
            n = n / 10
        }
        if digits % 2 == 0 {
            output += 1
        }
    }
    return output
}
