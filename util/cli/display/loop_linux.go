//go:build linux
// +build linux

package display

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/emrebicer/sudoku-solver/util"
)

func Loop(newBoard [9][9]int) {

	// Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		fmt.Printf("Enter 1-9 to highlight numbers, 0 to exit\n")
		var input int

		os.Stdin.Read(b)
		input, err := strconv.Atoi(string(b))
		if err != nil {
			fmt.Printf("failed to parse the input with the error: %s, please only use number between 0-9...", err.Error())
			break
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
			fmt.Printf("Invalid input.\n")
		}
	}
}
