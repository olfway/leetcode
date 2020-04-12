class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:

        while len(stones) >= 2:
            stones.sort()
            y = stones.pop()
            x = stones.pop()
            if y == x:
                continue
            stones.append(y-x)

        if stones:
            return stones.pop()

        return 0

