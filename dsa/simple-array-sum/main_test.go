package main

import "testing"

func TestSimpleArraySum(t *testing.T) {
	got := simpleArraySum([]int32{1, 2, 3, 4, 10, 11})
	want := int32(31)

	if got != want {
		t.Fatalf("simpleArraySum() = %d, want %d", got, want)
	}
}
