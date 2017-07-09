package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4}

	sort := func(arr []int) {
		isDone := false
		size := len(arr)

		for !isDone {
			isSorted := true

			for i := 0; i < size-1; i++ {
				current := arr[i]
				next := arr[i+1]
				if current > next {
					// Swap 2 element of the array
					arr[i], arr[i+1] = next, current
					isSorted = false
				}
			}
			isDone = isSorted
		}
	}
	sort(originalArr)
	fmt.Println(originalArr)
}
