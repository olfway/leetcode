class Solution:

    def canJump(self, nums: List[int]) -> bool:

        last = len(nums) - 1
        first = last

        for index in range(last, -1, -1):

            if index + nums[index] >= first:
                first = min(first, index)

        return first == 0

