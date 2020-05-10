class Solution:

    def findJudge(self, N: int, trust: List[List[int]]) -> int:

        if N == 1:
            return 1

        trusts = [0] * N
        for person, trusted in trust:
            trusts[person-1] -= 1
            trusts[trusted-1] +=1

        for n, v in enumerate(trusts):
            if v == N - 1:
                return n+1

        return -1

