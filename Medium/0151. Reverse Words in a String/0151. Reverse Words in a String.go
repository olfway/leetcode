import (
    "strings"
)

func reverseWords(s string) string {

    a, b := len(s), 0
    var result strings.Builder

    for {
        for b = a - 1; b >= 0 && s[b] == ' '; b-- {
        }

        if b < 0 { break }

        for a = b; a >= 0 && s[a] != ' '; a-- {
        }

        if result.Len() > 0 { result.WriteByte(' ') }

        result.WriteString(s[a+1:b+1])
    }
    return result.String()
}
