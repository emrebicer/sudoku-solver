package main

import (
	"flag"
	"fmt"

	"github.com/emrebicer/sudoku-solver/solver"
	"github.com/emrebicer/sudoku-solver/util"
	"github.com/emrebicer/sudoku-solver/util/cli/display"
)

func main() {
	filename := flag.String("board", "", "The file path which represents the board, consist of 9 rows where each row has number from 0 to 9, 0 represents an empty cell")
	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		return
	}

	board, err := util.ReadSudokuFromFile(*filename)

	if err != nil {
		fmt.Printf("faile to read sudoku from file: %s", err)
		panic(err)
	}

	res, valid := solver.SolveBoard(board)

	if !valid {
		fmt.Printf("Could not solve...\n")
		return
	}

	util.PrintBoard(res, 0)
	display.Loop(res)
}
