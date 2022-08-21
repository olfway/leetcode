func titleToNumber(columnTitle string) int {
    var n int
    for i := 0; i < len(columnTitle); i++ {
        n = n * 26 + int(columnTitle[i] - 'A') + 1
    }
    return n
}

