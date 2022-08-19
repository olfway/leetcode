func generate(numRows int) [][]int {
    var t [][]int
    var row, col int
    for row = 0; row < numRows; row++ {
        t = append(t, make([]int, row+1, row+1))
        t[row][0] = 1
        t[row][row] = 1
        for col = 1; col < row; col++ {
            t[row][col] = t[row-1][col-1] + t[row-1][col]
        }
    }
    return t
}
