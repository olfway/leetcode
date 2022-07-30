func replaceElements(arr []int) []int {
    max := -1
    var tmp int
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] > max {
            tmp = arr[i]
            arr[i] = max
            max = tmp
        } else {
            arr[i] = max
        }
    }
    return arr
}
