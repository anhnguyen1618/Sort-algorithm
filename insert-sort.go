package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4}

	insertSort := func(arr []int) {
		size := len(arr)

		swap := func(arr []int, firstIndex int, secondIndex int) {
			arr[firstIndex], arr[secondIndex] = arr[secondIndex], arr[firstIndex]
		}

		findSmallestIndex := func(arr []int, currentIndex int) int {
			smallest, smallIndex := arr[currentIndex], currentIndex
			for i := currentIndex + 1; i < size; i++ {
				if arr[i] < smallest {
					smallest = arr[i]
					smallIndex = i
				}
			}
			return smallIndex
		}

		for i := 0; i < size; i++ {
			swap(arr, i, findSmallestIndex(arr, i))
		}
	}

	insertSort(originalArr)
	fmt.Println(originalArr)
}
