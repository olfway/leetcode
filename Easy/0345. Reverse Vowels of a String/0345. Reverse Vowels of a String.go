import (
    "unsafe"
)

var vowels = map[byte]bool {
    'A': true, 'a': true,
    'E': true, 'e': true,
    'I': true, 'i': true,
    'O': true, 'o': true,
    'U': true, 'u': true,
}

func reverseVowels(s string) string {
    length := len(s)
    source := unsafe.Slice(unsafe.StringData(s), length)
    result := make([]byte, length)
    copy(result, source)

    b := length

    for a := 0; a < length; a++ {
        if vowels[result[a]] == false {
            continue
        }
        b -= 1
        for vowels[result[b]] == false {
            b--
        }
        if a >= b {
            break
        }
        result[a], result[b] = result[b], result[a]
    }

    return unsafe.String(unsafe.SliceData(result), len(result))
}
