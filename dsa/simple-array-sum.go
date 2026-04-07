package main

import (
	"fmt"
)

func simpleArraySum(ar []int32) int32 {
	var sum int32

	for _, v := range ar {
		sum += v
	}

	return sum
}

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	result := simpleArraySum(ar)
	fmt.Println(result)
}
