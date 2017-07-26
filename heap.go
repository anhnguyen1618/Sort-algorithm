package main

import "fmt"

func main() {
	originalArr := []int{1, 3, 12, 2, 6, 8, 5, 4}
	heapSort(originalArr)
	fmt.Println(originalArr)

}

func heapSort(arr []int) {
	size := len(arr)
	createMaxHeapInRange(arr, size)
	arr[0], arr[size-1] = arr[size-1], arr[0]

	for size = size - 1; size > 1; size-- {
		maxHeapify(arr, 0, size)

		/*
		*	Swap first and last element of the remaining heap everytime we decrement the size
		* of the heap so that the first node of the heap is excluded from the next heap
		 */

		arr[0], arr[size-1] = arr[size-1], arr[0]
	}
}

func createMaxHeapInRange(arr []int, size int) {
	for i := int((size - 1) / 2); i >= 0; i-- {
		maxHeapify(arr, i, size)
	}
}

func maxHeapify(arr []int, index int, size int) {

	leftIndex, rightIndex := index*2+1, index*2+2

	largest, indexLargest := arr[index], index

	if rightIndex < size && arr[rightIndex] > largest {
		largest, indexLargest = arr[rightIndex], rightIndex
	}

	if leftIndex < size && arr[leftIndex] > largest {
		largest, indexLargest = arr[leftIndex], leftIndex
	}

	if largest != arr[index] {
		arr[index], arr[indexLargest] = largest, arr[index]
		maxHeapify(arr, indexLargest, size)
	}

}
