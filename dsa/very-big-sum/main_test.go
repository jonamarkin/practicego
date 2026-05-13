package main

import "testing"

func TestAVeryBigSum(t *testing.T) {
	got := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	want := int64(5000000015)

	if got != want {
		t.Fatalf("aVeryBigSum() = %d, want %d", got, want)
	}
}
