# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:

    def firstBadVersion(self, n):

        a, b = 1, n

        while a != b:

            i = a + (b - a) // 2

            if isBadVersion(i):
                b = i
            else:
                a = max(i, a + 1)

        return a

