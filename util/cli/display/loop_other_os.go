//go:build !linux
// +build !linux

package display

import (
	"fmt"

	"github.com/emrebicer/sudoku-solver/util"
)

func Loop(newBoard [9][9]int) {

	for {

		fmt.Printf("Enter 1-9 to highlight numbers, 0 to exit\n")
		var input int

		_, inputErr := fmt.Scanln(&input)
		if inputErr != nil {
			fmt.Printf(inputErr.Error())
		}

		if input == 0 {
			break
		} else if input > 0 && input < 10 {
			// Clear the screen
			for i := 0; i < 50; i++ {
				fmt.Println()
			}
			util.PrintBoard(newBoard, input)
		} else {
			fmt.Println("Invalid input, please only use number from 0 to 9")
		}

	}
}
