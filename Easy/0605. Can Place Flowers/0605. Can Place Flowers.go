func canPlaceFlowers(flowerbed []int, n int) bool {

    if n == 0 {
        return true
    }

    length := len(flowerbed)

    if length == 1 {
        if n > 1 {
            return false
        }
        if flowerbed[0] == 1 {
            return false
        }
        return true
    }

    if length > 1 {
        if flowerbed[0] == 0 && flowerbed[1] == 0 {
            flowerbed[0] = 1
            n -= 1
        }

        if flowerbed[length-2] == 0 && flowerbed[length-1] == 0 {
            flowerbed[length-1] = 1
            n -= 1
        }
    }

    if n < 1 {
        return true
    }

    for i := 1; i < length - 1; i++ {
        if flowerbed[i-1] == 1 {
            continue
        }
        if flowerbed[i] == 1 {
            i += 1
            continue
        }
        if flowerbed[i+1] == 1 {
            i += 2
            continue
        }
        flowerbed[i] = 1
        n -= 1
        if n == 0 {
            return true
        }
    }
    return false
}
