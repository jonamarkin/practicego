package main

func countLengthOfCycle(arr []int, startIndex int) int {
	n := len(arr)

	if startIndex < 0 || startIndex >= n {
		return -1
	}
	slow := startIndex

	if slow < 0 || slow >= n {
		return -1
	}
	fast := arr[slow]

	for slow != fast {

		//Move slow pointer by 1 step
		if slow < 0 || slow >= n {
			return -1
		}
		slow = arr[slow]

		//Move fast pointer by 2 steps
		if fast < 0 || fast >= n {
			return -1
		}
		fast = arr[fast]

		if fast < 0 || fast >= n {
			return -1
		}
		fast = arr[fast]
	}

	// Cycle detected so we need to find the length of the cycle
	cycleLength := 1
	current := arr[slow]

	for current != slow {
		current = arr[current]
		cycleLength++
	}
	return cycleLength
}
