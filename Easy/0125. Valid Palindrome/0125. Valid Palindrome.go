func alphaNum(b byte) byte {
    if b >= byte('A') && b <= byte('Z') {
        b = b - byte('A') + byte('a')
    }
    if b < byte('0') || b > byte('9') {
        if b < byte('a') || b > byte('z') {
            return 0
        }
    }
    return b
}

func isPalindrome(s string) bool {

    i := 0
    j := len(s) - 1
    var b1, b2 byte

    for i < j {
        b1 = alphaNum(byte(s[i]))
        if b1 == 0 {
            i++
            continue
        }

        b2 = alphaNum(byte(s[j]))
        if b2 == 0 {
            j--
            continue
        }

        if b1 != b2 {
            return false
        }
        i++
        j--
    }
    return true
}
