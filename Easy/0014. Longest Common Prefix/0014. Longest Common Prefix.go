func longestCommonPrefix(strs []string) string {

    prefix := ""

    for {
        for i := 0; i < len(strs); i++ {
            if i == 0 {
                if len(prefix) == len(strs[i]) {
                    return prefix
                }
                prefix = strs[i][:len(prefix)+1]
                continue
            }
            if len(prefix) > len(strs[i]) {
                return prefix[:len(prefix)-1]
            }
            if prefix[len(prefix)-1] != strs[i][len(prefix)-1] {
                return prefix[:len(prefix)-1]
            }
        }
    }
}
