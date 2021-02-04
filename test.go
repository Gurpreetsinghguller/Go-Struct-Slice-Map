package main

import "fmt"

func main() {
	nn := []int{1, 2, 3, 4, 5}
	for index, value := range nn {
		fmt.Println(index, value)
	}
}
