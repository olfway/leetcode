func kidsWithCandies(candies []int, extraCandies int) []bool {
    kids := len(candies)
    result := make([]bool, kids)

    maxCandies := 0
    for _, n := range candies {
        maxCandies = max(maxCandies, n)
    }

    for i, n := range candies {
        if n + extraCandies >= maxCandies {
            result[i] = true
        }
    }

    return result
}
