func validMountainArray(arr []int) bool {

    length := len(arr)

    if length < 3 {
        return false
    }

    var i, n int
    for i = 0 ; i < length - 1; i++ {
        n = arr[i] - arr[i+1]
        if n == 0 {
            return false
        }
        if n > 0 {
            break
        }
    }

    if i == 0 {
        return false
    }

    if i == length - 1 {
        return false
    }

    for i = i; i < length - 1; i++ {
        n = arr[i] - arr[i+1]
        if n <= 0 {
            return false
        }
    }
    return true
}
