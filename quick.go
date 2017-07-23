package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4, 100, 12, 100, 1, 1, 12, 100, 1, 12, 100, 1, 1}
	newArr := sort(originalArr)
	fmt.Println(newArr)
}

func sort(arr []int) []int {
	size := len(arr)
	if size == 0 {
		return []int{}
	}

	if size == 1 {
		return arr
	}

	pivot := arr[size-1]
	leftArr := []int{}
	rightArr := []int{}

	for i := 0; i < size-1; i++ {
		currentElement := arr[i]
		if currentElement < pivot {
			leftArr = append(leftArr, currentElement)
		} else {
			rightArr = append(rightArr, currentElement)
		}
	}

	return append(append(sort(leftArr), pivot), sort(rightArr)...)
}
