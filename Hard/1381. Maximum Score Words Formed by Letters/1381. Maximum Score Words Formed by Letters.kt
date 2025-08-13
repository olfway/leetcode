fun CharSequence.toLetterCounts(): IntArray {
    val counts = IntArray(26)
    for (char in this) {
        counts[char - 'a']++
    }
    return counts
}

class Solution {
    fun maxScoreWords(words: Array<String>, letters: CharArray, score: IntArray): Int {
        val letterCounts = letters.concatToString().toLetterCounts()

        val wordCounts = words.associateWith { word ->
            word.toLetterCounts()
        }

        val wordScores = words.associateWith {
            it.sumOf { char -> score[char - 'a'] }
        }

        val wordSets = words.fold(
            initial = mapOf(emptyList<String>() to letterCounts)
        ) { wordSets, word ->
            val newWordSets = wordSets.mapNotNull { (key, remainingCounts) ->
                val newRemainingCounts = IntArray(26) { i ->
                    remainingCounts[i] - wordCounts.getValue(word)[i]
                }
                if (newRemainingCounts.all { it >= 0 }) {
                    (key + word) to newRemainingCounts
                } else {
                    null
                }
            }
            wordSets + newWordSets
        }

        return wordSets.map {
            it.key.sumOf { word -> wordScores.getValue(word) }
        }.max()
    }
}
