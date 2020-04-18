class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:

        if not grid:
            return 0

        rows = len(grid)
        cols = len(grid[0])

        def markNeighbors(x, y):
            q = list([(x, y)])
            while q:
                x, y = q.pop()
                for i, j in (
                    (x-1, y), (x+1, y),
                    (x, y-1), (x, y+1)
                ):
                    if 0 <= i < rows and 0 <= j < cols:
                        if grid[i][j] == "1":
                            grid[i][j] = '*'
                            q.append((i, j))

        islands = 0
        for x in range(rows):
            for y in range(cols):
                if grid[x][y] == "1":
                    islands += 1
                    grid[x][y] = '*'
                    markNeighbors(x, y)

        return islands
