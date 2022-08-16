func mySqrt(x int) int {
    n := 0
    a, b := 0, 46341
    if x == 0 {
        return 0
    }
    if b * b == x {
        return b
    }
    for {
        if b - a <= 1 {
            return a
        }
        n = a + (b - a) / 2
        if n == a {
            n++
        }
        if n * n == x {
            return n
        }
        if n * n < x {
            a = n
            continue
        }
        b = n
    }
}
