package solver

const emptyCell = 0

// Solve the given board
// The empty cells on the board should be indicated with a 0
func SolveBoard(board [9][9]int) ([9][9]int, bool) {

	// Find next empty cell
	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == emptyCell {
				validFlag := false
				for k := 1; k < 10; k++ {
					board[i][j] = k
					cellValid := isCellValid(&board, i, j)
					if cellValid {
						newBoard, boardValid := SolveBoard(board)
						if boardValid {
							validFlag = true
							board = newBoard
							break
						}
					}
				}

				if !validFlag {
					return board, false
				}

			}
		}
	}
	return board, true
}

func isCellValid(board *[9][9]int, rowIndex int, columnIndex int) bool {

	cellNumber := board[rowIndex][columnIndex]
	// Check if the row consists of the same number
	foundAtRowCounter := 0
	for i := range len(board[rowIndex]) {
		if board[rowIndex][i] == cellNumber {
			foundAtRowCounter++
		}
	}

	if foundAtRowCounter > 1 {
		return false
	}

	// Check if the column consists of the same number
	foundAtColumnCounter := 0
	for i := range len(board) {
		if board[i][columnIndex] == cellNumber {
			foundAtColumnCounter++
		}
	}

	if foundAtColumnCounter > 1 {
		return false
	}

	// Check if the current big square consist of the same number
	startRow := rowIndex - rowIndex%3
	startCol := columnIndex - columnIndex%3
	foundAtBigSquare := 0
	for i := range 3 {
		for j := range 3 {
			if board[i+startRow][j+startCol] == cellNumber {
				foundAtBigSquare++
			}
		}
	}

	if foundAtBigSquare > 1 {
		return false
	}

	return true
}
