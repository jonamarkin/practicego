package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var leftDiagonalSum, rightDiagonalSum int32

	for i := 0; i < len(arr); i++ {
		leftDiagonalSum += arr[i][i]
		rightDiagonalSum += arr[i][len(arr)-1-i]
	}

	diff := leftDiagonalSum - rightDiagonalSum
	if diff < 0 {
		return -diff
	}

	return diff
}

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	fmt.Println(diagonalDifference(arr))
}
