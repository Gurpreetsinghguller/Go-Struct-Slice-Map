//function takes 2d array of any size....return its transpose.
//Function k andar size static nai bnana hai
package main

import "fmt"

func Transpose(sli [][]int) [][]int {

	row := len(sli)
	col := len(sli[0])
	temp := make([][]int, row)
	// fmt.Println(temp)

	for i := 0; i < row; i++ {
		temp[i] = make([]int, col)

	}
	// fmt.Println(temp)

	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli[i]); j++ {
			temp[i][j] = sli[i][j]
		}

	}
	// fmt.Println(temp)
	for i := 0; i < row; i++ {
		for j := 0; j < len(sli); j++ {
			sli[i][j] = temp[j][i]
		}

	}
	return sli
}

func main() {
	var sli [][]int = [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 2},
		{7, 8, 9, 3},
		{10, 11, 12, 4},
	}
	fmt.Println(sli)
	fmt.Println(Transpose(sli))

}
