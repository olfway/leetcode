func checkIfExist(arr []int) bool {
    h := make(map[int]struct{}, len(arr))
    for i := 0; i < len(arr); i++ {
        n := arr[i]
        if _, ok := h[2 * n]; ok {
            return true
        }
        if n % 2 == 0 {
            if _, ok := h[n / 2]; ok {
                return true
            }
        }
        h[n] = struct{}{}
    }
    return false
}
