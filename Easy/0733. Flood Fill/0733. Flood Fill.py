class Solution:

    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:

        rows = len(image)
        cols = len(image[0])
        color = image[sr][sc]

        queue = {(sr, sc)}
        visited = {(sr, sc)}

        def queuePoint(row, col):

            if (0 <= row < rows) and (0 <= col < cols):
                point = (row, col)
                if point not in visited:
                    visited.add(point)
                    if image[row][col] == color:
                        queue.add(point)

        while queue:

            row, col = queue.pop()

            image[row][col] = newColor

            queuePoint(row-1, col)
            queuePoint(row+1, col)
            queuePoint(row, col-1)
            queuePoint(row, col+1)

        return image

