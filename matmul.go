package main

import "fmt"

func Multiplication(matrixA, matrixB [][]int, row1, col1, row2, col2 int) [][]int {
	var total int
	result := make([][]int, row1)
	for i := 0; i < row1; i++ {
		result[i] = make([]int, col2)
	}
	for i := 0; i < row1; i++ {
		for j := 0; j < col2; j++ {
			for k := 0; k < row2; k++ {
				total = total + matrixA[i][k]*matrixB[k][j]
			}
			result[i][j] = total
			total = 0
		}
	}
	return result
}
func main() {

	var row1, col1, row2, col2 int

	fmt.Println("Enter number of rows for Matrix 1")
	fmt.Scan(&row1)

	fmt.Println("Enter number of columns for Matrix 1 ")
	fmt.Scan(&col1)

	matrixA := make([][]int, row1)

	for i := 0; i < row1; i++ {
		matrixA[i] = make([]int, col1)
	}

	fmt.Println("Enter number of rows for Matrix 2")
	fmt.Scan(&row2)

	fmt.Println("Enter number of columns for Matrix 2")
	fmt.Scan(&col2)

	matrixB := make([][]int, row2)

	for i := 0; i < row2; i++ {
		matrixB[i] = make([]int, col2)
	}
	if col1 != row2 {
		fmt.Println("Multiplication cannot be performed")
	} else {
		fmt.Println("Enter the first matrix elements: ")

		for i := 0; i < row1; i++ {
			for j := 0; j < col1; j++ {
				fmt.Scan(&matrixA[i][j])
			}
		}
		fmt.Println("Enter the second matrix elements: ")
		for i := 0; i < row2; i++ {
			for j := 0; j < col2; j++ {
				fmt.Scan(&matrixB[i][j])
			}
		}
	}
	fmt.Println(Multiplication(matrixA, matrixB, row1, col1, row2, col2))

}
