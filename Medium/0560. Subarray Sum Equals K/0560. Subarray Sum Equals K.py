class Solution:

    def subarraySum(self, nums: List[int], k: int) -> int:

        sum = 0
        count = 0
        sums = {0: 1}

        for num in nums:
            sum += num
            count += sums.get(sum - k, 0)
            sums[sum] = sums.get(sum, 0) + 1

        return count

