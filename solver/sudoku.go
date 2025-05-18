package solver

import (
	"github.com/emrebicer/sudoku-solver/util"
)

const emptyCell = 0

// Solve the given board
// The empty cells on the board should be indicated with a 0
func SolveBoard(board *[9][9]int) bool {
	// Find next empty cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == emptyCell {
				validFlag := false
				for k := 1; k < 10; k++ {
					board[i][j] = k
					cellValid := isCellValid(board, i, j)
					if cellValid {
						// Copy the board
						auxBoard := [9][9]int{}
						util.CopyBoard(&auxBoard, board)
						boardValid := SolveBoard(&auxBoard)
						if boardValid {
							validFlag = true
							util.CopyBoard(board, &auxBoard)
							break
						}
					}
				}

				if !validFlag {
					return false
				}

			}
		}
	}
	return true
}

func isCellValid(board *[9][9]int, rowIndex int, columnIndex int) bool {

	cellNumber := board[rowIndex][columnIndex]
	// Check if the row consists of the same number
	foundAtRowCounter := 0
	for i := 0; i < len(board[rowIndex]); i++ {
		if board[rowIndex][i] == cellNumber {
			foundAtRowCounter++
		}
	}

	if foundAtRowCounter > 1 {
		return false
	}

	// Check if the column consists of the same number
	foundAtColumnCounter := 0
	for i := 0; i < len(board); i++ {
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
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
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
