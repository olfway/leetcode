class Solution:

    def countElements(self, arr: List[int]) -> int:

        r = 0
        counters = dict()

        for num in arr:

            if num not in counters:
                if (num - 1) in counters:
                    r += counters[num - 1]

            counters[num] = counters.get(num, 0) + 1

            if (num + 1) in counters:
                r += 1

        return r

