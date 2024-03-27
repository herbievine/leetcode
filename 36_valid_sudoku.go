package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	grid := make([]map[byte]bool, 3)

	for i := range board {
		row := make(map[byte]bool)
		col := make(map[byte]bool)

		for j := range board[i] {
			gridIndex := (j/3 + i/3) % 3

			if i%3 == 0 && j%3 == 0 {
				grid[gridIndex] = make(map[byte]bool)
			}

			if board[i][j] != '.' {
				if _, ok := grid[gridIndex][board[i][j]]; ok {
					return false
				}
				grid[gridIndex][board[i][j]] = true
			}

			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}
		}
	}

	return true
}

func main() {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(sudoku))
}
