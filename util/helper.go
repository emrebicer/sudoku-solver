package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func ReadSudokuFromFile(filename string) ([9][9]int, error) {

	board := [9][9]int{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return board, err
	}

	lines := strings.Split(string(data), "\n")

	for i := range 9 {
		currentLine := lines[i]
		for j := range 9 {
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
	highlightColor := color.FgGreen

	color.Set(color.FgWhite)
	width := 37
	third := int(width / 3)
	for i := range width {
		if i == 0 || i == width-1 || i%third == 0 {
			fmt.Printf("+")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Println()

	for i := range 9 {
		fmt.Printf("|")
		for j := range 9 {
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
		for j := range width {
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
