func duplicateZeros(arr []int)  {
    max := len(arr) - 1
    for i := 0; i < max; i++ {
        if arr[i] != 0 {
            continue
        }
        for j := max ; j - 1 > i; j-- {
            arr[j] = arr[j-1]
        }
        i++
        arr[i] = 0
    }
}
