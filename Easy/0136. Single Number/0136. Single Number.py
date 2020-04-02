class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        buf = set()
        for num in nums:
            if num in buf:
                buf.remove(num)
                continue
            buf.add(num)

        return list(buf)[0]

