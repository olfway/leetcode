class Solution:
    def rangeBitwiseAnd(self, m: int, n: int) -> int:

        bits = 0
        while True:

            if n == m:
                return m << bits

            bits += 1
            m = m >> 1
            n = n >> 1

