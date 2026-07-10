package countlengthcycle

func countLengtOfCycle(arr []int, startIndex int) int {
	visitedSteps := make(map[int]int)

	currentIndex := startIndex

	steps := 0

	for currentIndex >= 0 && currentIndex < len(arr) {

		if firstSeenStep, exists := visitedSteps[currentIndex]; exists {
			return steps - firstSeenStep
		}

		visitedSteps[currentIndex] = steps

		currentIndex = arr[currentIndex]
		steps++
	}

	return -1
}

// func countLengthOfCycleOptimal(arr []int, startIndex int) int {
// 	slow :
// }
