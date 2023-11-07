// Package valid_sudoku
// difficulty: easy
package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	// check columns and rows
	for i := 0; i < 9; i++ {
		memCol := make(map[byte]struct{})
		memRow := make(map[byte]struct{})
		memSquare := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			colCell := board[j][i]
			rowCell := board[i][j]
			squareX, squareY := toSquareIndexes(i, j)
			squareCell := board[squareX][squareY]
			if !isInMem(colCell, &memCol) || !isInMem(rowCell, &memRow) || !isInMem(squareCell, &memSquare) {
				return false
			}
		}
	}

	// check squares

	return true

}

func isInMem(char byte, mem *map[byte]struct{}) bool {
	if char != '.' {
		if !charIsValid(char) {
			return false
		}
		if _, exists := (*mem)[char]; exists {
			return false
		}
		(*mem)[char] = struct{}{}
	}
	return true
}

func charIsValid(char byte) bool {
	if char == '.' {
		return true
	}
	if char-'0' > 0 && char-'0' < 10 {
		return true
	}
	return false
}

func toSquareIndexes(i, j int) (int, int) {
	y := (j / 3) + 3*(i/3)
	x := (j % 3) + 3*(i%3)
	return x, y
}
