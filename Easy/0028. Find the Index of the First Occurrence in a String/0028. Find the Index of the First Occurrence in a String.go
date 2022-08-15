func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    if len(needle) > len(haystack) {
        return -1
    }
    var i, j int
    for i = 0; i < len(haystack) - len(needle) + 1; i++ {
        if haystack[i] != needle[0] {
            continue
        }
        for j = 0; j < len(needle); j++ {
            if haystack[i+j] != needle[j] {
                break
            }
        }
        if j == len(needle) {
            return i
        }
    }
    return -1
}
