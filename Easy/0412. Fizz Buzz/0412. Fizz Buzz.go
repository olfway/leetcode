import "strconv"

func fizzBuzz(n int) []string {
    output := make([]string, 0, n)
    for i := 1; i <= n; i++ {
        switch {
        case i % 3 == 0 && i % 5 == 0:
            output = append(output, "FizzBuzz")
            continue
        case i % 3 == 0:
            output = append(output, "Fizz")
            continue
        case i % 5 == 0:
            output = append(output, "Buzz")
            continue
        default:
            output = append(output, strconv.Itoa(i))
        }
    }
    return output
}
