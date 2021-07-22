package main

import (
	"fmt"
)

func main() {

	var board = []int{1, 2, 3, 4, 5, 16}

	for i := range board {
		res := board[i] << 1
		isOdd := board[i] & 1
		fmt.Printf("moving %d one bit to the left. result: %d\n", board[i], res)
		fmt.Printf("%d is even: %v", board[i], isEven(b))
	}
}

// square moves a bit to the left using the bitwise operator.
func square(num int) int {
	return num << 1
}

// is even checks if 1 in binary format is present 
func isEven(num int) bool {
	res := num & 1

	if res == 0 {
		return true
	}

	return false
}

