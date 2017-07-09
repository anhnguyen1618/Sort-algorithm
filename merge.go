package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4}

	originalArr = sort(originalArr)
	fmt.Println(originalArr)
}

func sort(arr []int) []int {
	size := len(arr)
	if size == 0 {
		return []int{}
	}

	if size == 1 {
		return arr
	}

	middleIndex := int(size / 2)
	leftArr := sort(arr[:middleIndex])
	rightArr := sort(arr[middleIndex:])

	return merge(leftArr, rightArr)
}

/*
	Merge 2 sorted arrays in ascending order
*/
func merge(leftArr []int, rightArr []int) []int {
	temp := []int{}
	leftLength, rightLength := len(leftArr), len(rightArr)
	i, j := 0, 0

	for i < leftLength && j < rightLength {
		if leftArr[i] < rightArr[j] {
			temp = append(temp, leftArr[i])
			i++
		} else {
			temp = append(temp, rightArr[j])
			j++
		}
	}

	if i < leftLength {
		temp = append(temp, leftArr[i:]...)
	}

	if j < rightLength {
		temp = append(temp, rightArr[j:]...)
	}

	return temp
}
