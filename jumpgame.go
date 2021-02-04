package main

import "fmt"

func jump(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}

	currentPosition := arr[0]
	maxJump := arr[0] + 0
	jumpCount := 1
	for i := 1; i < len(arr); i++ {
		// newPosition := arr[i]
		newJump := arr[i] + i
		if newJump > maxJump {
			maxJump = newJump
		}
		if i == currentPosition {
			jumpCount += 1
		}
		currentPosition = maxJump

	}
	return jumpCount
}

func main() {
	arr := []int{1, 3, 6, 3, 2, 3, 6, 8, 9, 5}
	len := len(arr)
	fmt.Println(jumps(arr))

}
