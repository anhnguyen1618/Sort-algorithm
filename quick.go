package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 7, 6, 8, 1, 4, 5}
	quickSort(originalArr)
	fmt.Println(originalArr)
}

func partition(arr []int, start int, end int) int {
	wallIndex := start
	pivot := arr[end]
	for i := start; i < end; i++ {
		currentElement := arr[i]
		if currentElement < pivot {
			arr[wallIndex], arr[i] = currentElement, arr[wallIndex]
			wallIndex++
		}
	}

	arr[wallIndex], arr[end] = pivot, arr[wallIndex]
	return wallIndex
}

func sort(arr []int, start int, end int) {
	if start < end {
		pivotIndex := partition(arr, start, end)

		sort(arr, start, pivotIndex-1)
		sort(arr, pivotIndex+1, end)
	}
}

func quickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}
