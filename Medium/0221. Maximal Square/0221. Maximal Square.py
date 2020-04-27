class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:

        if not matrix:
            return 0

        rows = len(matrix)
        cols = len(matrix[0])

        res = 0
        state = [[0] * cols for _ in range(rows)]

        for row in range(0, rows):
            for col in range(0, cols):

                if matrix[row][col] != "1":
                    continue

                state[row][col] = 1 + min(
                    state[row-1][col] if row > 0 else 0,
                    state[row][col-1] if col > 0 else 0,
                    state[row-1][col-1] if min(row, col) > 0 else 0
                )
                res = max(res, state[row][col])

        return res*res

