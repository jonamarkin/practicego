package main

import "testing"

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	got := diagonalDifference(arr)
	want := int32(15)

	if got != want {
		t.Fatalf("diagonalDifference() = %d, want %d", got, want)
	}
}
