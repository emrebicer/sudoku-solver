package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/emrebicer/sudoku-solver/solver"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/solve", func(c *gin.Context) {
		boardStr := c.Request.URL.Query().Get("board")
		res, err := solveFromString(boardStr)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to solve the board: %s", err)
		}

		c.String(http.StatusOK, res)
	})

	router.Run("localhost:8080")
}

func solveFromString(boardStr string) (string, error) {

	// The board string must consist of 81 characters
	if len(boardStr) != 81 {
		return "", errors.New("board must consist of 81 elements")
	}

	// Each characters must only be digits
	if match, _ := regexp.MatchString(`^[0-9]+$`, boardStr); !match {
		return "", errors.New("board must only contain numbers")
	}

	// Construct the board
	var board [9][9]int
	for i := range 9 {
		for j := range 9 {
			index := i*9 + j
			val, err := strconv.Atoi(string(boardStr[index]))
			if err != nil {
				return "", fmt.Errorf("failed to construct integer from index: %d, %s", index, err)
			}

			board[i][j] = val
		}
	}
	solvedBoard, res := solver.SolveBoard(board)

	if !res {
		return "", fmt.Errorf("failed to solve the board")
	}

	result := ""
	for i := range 9 {
		for j := range 9 {
			result += strconv.Itoa(solvedBoard[i][j])
		}
	}

	return result, nil
}
