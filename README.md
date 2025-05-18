# Sudoku Solver
A brute force 9x9 sudoku solver written in Go
## Use the CLI solver
Create a file that represents the sudoku board, use 0 for empty cells. E.g.
```
080009005
000310007
000005403
406000020
190000076
030000904
308700000
700063000
200800030
```
Run `cli_solver` example to solve the sudoku and display the result on your terminal
```terminal
cd cmd/cli_solver
go run main.go --board path/to/board.txt
```
## Use the solver module
Import the project and use the solve.SolveBoard method, your argument will be mutated and will represent the solution if return argument is true.
```go
package main

import (
	"github.com/emrebicer/sudoku-solver/solver"
)

func main() {
	valid := solver.SolveBoard(&board)

	if !valid {
		fmt.Printf("Could not solve...\n")
		return
	}
}
```
