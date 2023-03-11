package main

import (
	"fmt"
)

func diagonalDifference(arr [][]int32) int32 {
	//write your code here
	var kiriDiagonal, kananDiagonal int32
	for i := 0; i < len(arr); i++ {
		kiriDiagonal += arr[i][i]
		kananDiagonal += arr[i][len(arr)-1-i]
	}
	return (kiriDiagonal - kananDiagonal)
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(diagonalDifference(arr)) // Output: 0
}