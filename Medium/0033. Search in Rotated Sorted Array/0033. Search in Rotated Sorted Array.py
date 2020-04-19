class Solution:
    def search(self, nums: List[int], target: int) -> int:

        left = 0
        right = len(nums) - 1

        if right == -1:
            return -1

        while left != right:

            middle = (left + right) // 2

            if nums[left] <= nums[middle]:
                if nums[left] <= target <= nums[middle]:
                    right = middle
                    continue

                left = middle if left != middle else left + 1
                continue

            if nums[middle] <= target <= nums[right]:
                left = middle
                continue

            right = middle if right != middle else right - 1

        if nums[left] == target:
            return left

        return -1

