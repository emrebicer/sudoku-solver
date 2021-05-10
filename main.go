package main

import (
	"flag"
	"fmt"

	sudoku "github.com/emrebicer/sudoku-solver/sudoku"
)

func main() {
	filename := flag.String("f", "", "The file name")
	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		return
	}

	board, err := sudoku.ReadSudokuFromFile(*filename)

	if err != nil {
		panic(err)
	}

	newBoard, valid := sudoku.SolveBoard(board)

	if !valid {
		fmt.Printf("Could not solve...\n")
		return
	}

	sudoku.PrintBoard(newBoard, 0)
	sudoku.Loop(newBoard)
}
