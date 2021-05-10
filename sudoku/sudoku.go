package sudoku

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const emptyCell = 0

func SolveBoard(board [9][9]int) ([9][9]int, bool) {

	// Find next empty cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == emptyCell {
				validFlag := false
				for k := 1; k < 10; k++ {
					board[i][j] = k
					cellValid := isCellValid(board, i, j)
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

func isCellValid(board [9][9]int, rowIndex int, columnIndex int) bool {

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

func ReadSudokuFromFile(filename string) ([9][9]int, error) {

	board := [9][9]int{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return board, err
	}

	lines := strings.Split(string(data), "\n")

	for i := 0; i < 9; i++ {
		currentLine := lines[i]
		for j := 0; j < 9; j++ {
			currentNum, err := strconv.Atoi(string(currentLine[j]))
			if err != nil {
				return board, err
			}
			board[i][j] = currentNum
		}
	}

	return board, nil
}

func PrintBoard(board [9][9]int, highlight_number int) {

	boardDashColor := color.FgHiMagenta
	highlightColor := color.FgCyan

	width := 37
	third := int(width / 3)
	for i := 0; i < width; i++ {
		if i == 0 || i == width-1 || i%third == 0 {
			fmt.Printf("+")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Println()

	for i := 0; i < 9; i++ {
		fmt.Printf("|")
		for j := 0; j < 9; j++ {
			currentNumber := board[i][j]
			if currentNumber == highlight_number {
				color.Set(highlightColor)
				fmt.Printf(" %d ", board[i][j])
				color.Set(color.FgWhite)
			} else {
				fmt.Printf(" %d ", board[i][j])
			}
			if j == 2 || j == 5 {
				color.Set(boardDashColor)
				fmt.Printf("|")
				color.Set(color.FgWhite)
			} else {
				fmt.Printf("|")
			}
		}
		fmt.Println()
		for j := 0; j < width; j++ {
			if ((i+1)%3 == 0) && (j == 0 || j == width-1 || j%third == 0) {
				fmt.Printf("+")
			} else {
				if i == 2 || i == 5 {
					color.Set(boardDashColor)
					fmt.Printf("-")
					color.Set(color.FgWhite)
				} else {
					fmt.Printf("-")
				}
			}
		}
		fmt.Println()
	}
}
