class Solution:

    def removeKdigits(self, num: str, k: int) -> str:

        if k == 0:
            return num

        if k == len(num):
            return "0"

        start = 0
        for step in range(k):

            for index in range(start, len(num)-1):
                if num[index] > num[index+1]:
                    start = max(0, index-1)
                    num = num[:index] + num[index+1:]
                    break
            else:
                num = num[:-1*(k-step)]
                return num.lstrip("0") or "0"

        return num.lstrip("0") or "0"

