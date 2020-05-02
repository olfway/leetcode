class Solution:

    def numJewelsInStones(self, J: str, S: str) -> int:

        n = 0
        j = set(J)
        for s in S:
            if s in j:
                n += 1

        return n

