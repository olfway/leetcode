class Solution:

    def findMaxLength(self, nums: List[int]) -> int:

        sum = 0
        sums = {0: -1}
        max_length = 0

        for i in range(len(nums)):

            sum += -1 if nums[i] == 0 else 1

            if sum not in sums:
                sums[sum] = i
                continue

            max_length = max(max_length, i - sums[sum])

        return max_length

