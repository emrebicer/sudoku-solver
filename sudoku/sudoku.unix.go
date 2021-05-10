// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package sudoku

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
			fmt.Printf(err.Error())
		}

		if input == 0 {
			break
		} else if input > 0 && input < 10 {
			// Clear the screen
			for i := 0; i < 50; i++ {
				fmt.Println()
			}
			printBoard(newBoard, input)
		} else {
			fmt.Printf("Invalid input.\n")
		}
	}
}
