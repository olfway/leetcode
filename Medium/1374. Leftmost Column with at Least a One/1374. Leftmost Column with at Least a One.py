# """
# This is BinaryMatrix's API interface.
# You should not implement it, or speculate about its implementation
# """
#class BinaryMatrix(object):
#    def get(self, x: int, y: int) -> int:
#    def dimensions(self) -> list[]:

class Solution:
    def leftMostColumnWithOne(self, binaryMatrix: 'BinaryMatrix') -> int:

        rows, cols = binaryMatrix.dimensions()

        result = -1
        x, y = 0, cols - 1
        while x < rows and y >= 0:
            if binaryMatrix.get(x, y) == 1:
                result = y
                y -= 1
                continue
            x += 1

        return result

