class Solution:

    def kadane(self, nums, invert=False):

        curSum = 0
        maxSum = float("-inf")

        for num in nums:

            if invert:
                num *= -1

            curSum = curSum + num

            maxSum = max(curSum, maxSum)

            if curSum < 0:
                curSum = 0

        return maxSum


    def maxSubarraySumCircular(self, A: List[int]) -> int:

        maxKadane = self.kadane(A)

        if maxKadane < 0:
            return maxKadane

        maxWrap = sum(A) + self.kadane(A, invert=True)

        return max(maxKadane, maxWrap)

