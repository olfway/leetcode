class Solution:
    def checkValidString(self, s: str) -> bool:

        low = 0
        high = 0

        for c in s:

            if c == '(':
                low += 1
                high += 1
            else:
                if c == ')':
                    if high == 0:
                        return False
                    high -= 1
                else:
                    high += 1
                low = max(0, low - 1)

        return low == 0

