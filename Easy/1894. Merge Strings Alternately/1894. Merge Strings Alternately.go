func mergeAlternately(word1 string, word2 string) string {

    slice1 := []byte(word1)
    slice2 := []byte(word2)

    result := make([]byte, 0, len(word1) + len(word2))

    min_len := min(len(word1), len(word2))

    for i := 0; i < min_len; i++ {
        result = append(result, slice1[i])
        result = append(result, slice2[i])
    }

    result = append(result, slice1[min_len:]...)
    result = append(result, slice2[min_len:]...)

    return string(result)
}
