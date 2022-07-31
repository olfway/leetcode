func romanToInt(s string) int {

    convert := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    length := len(s)
    var i, n, output int

    for i = 0; i < length - 1; i++ {
        n = convert[s[i]]
        if n < convert[s[i+1]] {
            output -= n
        } else {
            output += n
        }
    }

    output += convert[s[i]]

    return output
}
