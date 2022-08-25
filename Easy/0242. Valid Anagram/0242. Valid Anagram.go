func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    var letters [26]int

    for i := 0; i < len(s); i++ {
        letters[s[i]-'a']++
    }

    for i := 0; i < len(t); i++ {
        letters[t[i]-'a']--
        if letters[t[i]-'a'] < 0 {
            return false
        }
    }

    return true
}
