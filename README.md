# Sudoku Solver
 A 9x9 sudoku solver written in Go
## Install
Fetch the repo and required go packages
```terminal
git clone https://github.com/emrebicer/sudoku-solver.git
cd sudoku-solver
go get
```
## Usage
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
Run main.go to solve the sudoku
```terminal
go run main.go -f board.txt
```
