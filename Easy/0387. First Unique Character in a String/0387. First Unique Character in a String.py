class Solution:

    def firstUniqChar(self, s: str) -> int:

        counts = {}
        for i, c in enumerate(s):
            counts[c] = counts.get(c, 0) + 1

        for i, c in enumerate(s):
            if counts[c] == 1:
                return i

        return -1

