from math import isclose

class Solution:

    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:

        if len(coordinates) <= 2:
            return True

        def calcSlope(a, b):
            if a[0] == b[0]:
                return 0
            return (a[1]-b[1])/(a[0]-b[0])

        slope = calcSlope(coordinates[0], coordinates[1])
        for i in range(2, len(coordinates)):
            if isclose(calcSlope(coordinates[0], coordinates[i]), slope) == False:
                return False

        return True

