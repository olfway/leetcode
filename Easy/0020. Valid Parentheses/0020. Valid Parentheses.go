func isValid(s string) bool {
    length := len(s)

    if length % 2 == 1 {
        return false
    }

    pairs := map[byte]byte{
        '(': ')',
        '{': '}',
        '[': ']',
        ')': '(',
        '}': '{',
        ']': '[',
    }

    stack := make([]byte, 0, length / 2)

    var c byte

    for i := 0; i < length; i++ {
        c = s[i]
        switch c {
        case '(', '{', '[':
            if len(stack) == cap(stack) {
                return false
            }
            stack = append(stack, c)
        default:
            if len(stack) == 0 {
                return false
            }
            if stack[len(stack) - 1] != pairs[c] {
                return false
            }
            stack = stack[:len(stack) - 1]
        }
    }

    if len(stack) > 0 {
        return false
    }

    return true
}
