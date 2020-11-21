package main

import (
	"fmt"
	"strings"
)

func main() {
	// tic-tac board [also a matrix]
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	// players take turn
	board[0][1] = "X"
	board[1][2] = "O"
	board[2][1] = "x"
	board[1][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
