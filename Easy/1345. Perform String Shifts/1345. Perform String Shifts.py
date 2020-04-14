class Solution:

    def stringShift(self, s: str, shift: List[List[int]]) -> str:

        n = 0
        for direction, amount in shift:
            n += amount if direction else -1 * amount

        n = n % len(s)

        if n == 0:
            return s

        if n < 0:
            n += len(s)

        return s[-1 * n:] + s[:len(s) - n]

