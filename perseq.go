package main

import (
	"fmt"
	"strconv"
)

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact((n - 1))
}
func combination(arr []string, k, n int) string {
	stringCouter := 0
	str := make([]string, fact(n))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if i != j && j != k && k != i {
					str[stringCouter] = arr[i] + arr[j] + arr[k]
					stringCouter++
					fmt.Println(str)
				}

			}
		}
	}

	return str[k]
}

func main() {
	var n int
	fmt.Println("Enter number of elements")
	fmt.Scan(&n)
	arr := make([]string, n)
	// fmt.Println(len(arr), cap(arr))
	// a := 0 + 1
	// fmt.Println(strings(a))
	for i := 0; i < n; i++ {
		s := strconv.Itoa(i + 1)
		arr[i] = s
	}
	fmt.Println(arr)
	k := 0
	fmt.Println("Enter k")
	fmt.Scan(&k)
	fmt.Println(combination(arr, k-1, n))
}

//The set[1,2,3,4,...n] contains a total of n! unique permutation
//For n=3
///Permutations are
//123
//132
//213
//231
//312
//321
//Given n and k, return the kth permutation sequence
//Example :- input n=3, k=3  output 213
