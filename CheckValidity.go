package main

import (
	"fmt"
	"strconv"
)

func checkValidity(sudoku [][]int) bool {
	mappingElements := make(map[string]int)
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[0]); j++ {

			currentRow := strconv.Itoa(sudoku[i][j]) + " row " + strconv.Itoa(i)
			currentCol := strconv.Itoa(sudoku[i][j]) + " col " + strconv.Itoa(j)
			grid := strconv.Itoa(sudoku[i][j]) + " box " + strconv.Itoa((i/3)*3) + " " + strconv.Itoa((j/3)*3)
			// fmt.Println(currentRow, currentCol, grid)

			//checking if same value present in row
			if _, value := mappingElements[currentRow]; value {
				fmt.Println(currentRow)
				return false
			}

			//checking if same value present in col
			if _, value := mappingElements[currentCol]; value {
				fmt.Println(currentCol)
				return false
			}

			//checking if same value present in col
			if _, value := mappingElements[grid]; value {
				fmt.Println(grid)
				return false
			}

			mappingElements[currentRow] = sudoku[i][j]
			mappingElements[currentCol] = sudoku[i][j]
			mappingElements[grid] = sudoku[i][j]
			// if i == 0 && j == 0 {
			// 	fmt.Println(mappingElements)
			// } else {
			// 	break
			// }
		}
	}
	return true
}
func main() {
	sudoku1 := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	fmt.Println(checkValidity(sudoku1))
}
