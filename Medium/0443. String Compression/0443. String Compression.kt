class Solution {
    fun compress(chars: CharArray): Int {
        var groupSize = 0
        var groupIndex = 0
        var writeIndex = 1
        for (readIndex in 1 until chars.size + 1) {
            if (chars[groupIndex] == chars.getOrNull(readIndex)) {
                continue
            }
            groupSize = readIndex - groupIndex
            if (groupSize > 1) {
                groupSize.toString().forEach { digit ->
                    chars[writeIndex++] = digit
                }
            }
            if (readIndex < chars.size) {
                groupIndex = readIndex
                chars[writeIndex++] = chars[readIndex]
            }
        }
        return writeIndex
    }
}
