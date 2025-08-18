class Solution {
    fun isSubsequence(s: String, t: String): Boolean {
        var tIndex = 0
        for (sChar in s) {
            tIndex = t.indexOf(sChar, tIndex).takeIf { it != -1 } ?: return false
            tIndex += 1
        }
        return true
    }
}
