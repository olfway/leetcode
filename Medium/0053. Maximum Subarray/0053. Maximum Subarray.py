class Solution:
    def maxSubArray(self, nums: List[int]) -> int:

        s = nums[0]
        r = nums[0]

        for n in nums[1:]:

            if s > 0:
                s = s + n
            else:
                s = n

            if s > r:
                r = s

        return r

