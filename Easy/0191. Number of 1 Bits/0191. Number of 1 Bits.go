func hammingWeight(num uint32) int {
    var output int
    for num > 0 {
        output, num = output + int(num % 2), num >> 1
    }
    return output
}
