class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:

        cols = len(text1)
        rows = len(text2)

        res = [ [0] * cols for _ in range(rows) ]

        for row in range(rows):

            for col in range(cols):

                if text1[col] == text2[row]:
                    res[row][col] = 1 + res[row-1][col-1] if min(row, col) > 0 else 1
                    continue

                res[row][col] = max(
                    res[row][col-1] if col > 0 else 0,
                    res[row-1][col] if row > 0 else 0,
                )

        return res[-1][-1]

