class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        null = None
        for index in range(len(nums)):

            if nums[index] == 0:
                if null is None:
                    null = index
                continue

            if null is None:
                continue

            nums[null] = nums[index]
            null = null + 1
            nums[index] = 0
