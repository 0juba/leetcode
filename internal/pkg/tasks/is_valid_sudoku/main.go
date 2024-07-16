package is_valid_sudoku

const (
	boardSize = 9
	dot       = '.'
	zero      = '0'
)

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < boardSize; i++ {
		if !unique(getRow(board, i)) {
			return false
		}

		if !unique(getCol(board, i)) {
			return false
		}

		if !unique(getSector(board, i)) {
			return false
		}
	}

	return true
}

func getRow(board [][]byte, rowNum int) [boardSize]byte {
	row := [boardSize]byte{}
	for c := 0; c < len(board); c++ {
		row[c] = board[rowNum][c]
	}

	return row
}

func getCol(board [][]byte, colNum int) [boardSize]byte {
	col := [boardSize]byte{}
	for r := 0; r < len(board); r++ {
		col[r] = board[r][colNum]
	}

	return col
}

var sectors = [boardSize][2]int{
	{0, 0}, {0, 3}, {0, 6},
	{3, 0}, {3, 3}, {3, 6},
	{6, 0}, {6, 3}, {6, 6},
}

func getSector(board [][]byte, sectorNum int) [boardSize]byte {
	sector := sectors[sectorNum]
	cells := [boardSize]byte{}

	k := 0
	for i := sector[0]; i < sector[0]+3; i++ {
		for j := sector[1]; j < sector[1]+3; j++ {
			cells[k] = board[i][j]
			k++
		}
	}

	return cells
}

func unique(cells [boardSize]byte) bool {
	row := [boardSize]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, c := range cells {
		if c != dot {
			num := int(c - zero)
			row[num-1] -= num
			if row[num-1] < 0 {
				return false
			}
		}
	}

	return true
}
