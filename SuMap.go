package main

import "fmt"

func isValid(sudoku map[int][]int, row, col, value int) bool {
	//Checking column
	getRow := sudoku[row]
	for j := 0; j < len(getRow); j++ {
		if getRow[j] == value {
			return false
		}
	}

	//checking row
	for i := 0; i < len(sudoku); i++ {
		getCol := sudoku[i]
		if getCol[col] == value {
			return false
		}
	}

	//Checking for Grid
	startingRow := (row / 3) * 3
	startingCol := (col / 3) * 3
	for i := startingRow; i < startingRow+3; i++ {
		grid := sudoku[i]
		for j := startingCol; j < startingCol+3; j++ {

			if grid[j] == value {

				return false
			}
		}

	}
	return true

}
func sudoSolve(sudoku map[int][]int, row, col int) {

	if row == 9 {
		for i := 0; i < len(sudoku); i++ {
			fmt.Println(sudoku[i])
		}
		return
	}

	nextRow := 0
	nextCol := 0

	if col == 8 {
		nextRow = row + 1
		nextCol = 0
	} else {
		nextRow = row
		nextCol = col + 1
	}
	sliceRow := sudoku[row]
	if sliceRow[col] != 0 {
		sudoSolve(sudoku, nextRow, nextCol)
	} else {
		for k := 1; k <= 9; k++ {
			if isValid(sudoku, row, col, k) == true {
				sliceRow[col] = k
				sudoSolve(sudoku, nextRow, nextCol)
				sliceRow[col] = 0
			}
		}
	}
}
func main() {
	sudoku := map[int][]int{
		0: []int{3, 0, 6, 0, 0, 8, 4, 0, 0},
		1: []int{5, 2, 0, 0, 0, 0, 0, 0, 0},
		2: []int{0, 8, 7, 0, 0, 0, 0, 3, 1},
		3: []int{0, 0, 3, 0, 1, 0, 0, 8, 0},
		4: []int{9, 0, 0, 8, 6, 3, 0, 0, 5},
		5: []int{0, 5, 0, 0, 9, 0, 6, 0, 0},
		6: []int{1, 3, 0, 0, 0, 0, 2, 5, 0},
		7: []int{0, 0, 0, 0, 0, 0, 0, 7, 4},
		8: []int{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	sudoSolve(sudoku, 0, 0)

}
