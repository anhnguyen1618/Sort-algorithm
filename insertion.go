package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4}

	sort := func(arr []int) {
		size := len(arr)
		for i := 1; i < size; i++ {
			j := i - 1
			for j >= 0 && arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				j--
			}
		}
	}
	sort(originalArr)
	fmt.Println(originalArr)
}
