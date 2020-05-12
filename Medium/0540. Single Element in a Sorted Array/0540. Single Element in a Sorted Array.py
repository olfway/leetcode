class Solution:

    def singleNonDuplicate(self, nums: List[int]) -> int:

        low = 0
        high = len(nums) - 1

        while low < high:

            mid = low + (high - low) // 2

            if mid % 2 == 0:

                if nums[mid] == nums[mid + 1]:
                    low = mid + 2
                    continue

                high = mid
                continue

            if nums[mid] == nums[mid - 1]:
                low = mid + 1
                continue

            high = mid - 1

        return nums[low]

