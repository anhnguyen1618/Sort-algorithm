package main

import "fmt"

func main() {
	originalArr := []int{9, 3, 12, 2, 6, 8, 1, 4}

	sort := func(arr []int) {
		size := len(arr)
		for i := 1; i < size; i++ {
			for j := 0; j < i; j++ {
				if arr[i] < arr[j] {
					current := arr[i]

					//shift all element from index j => i by one position
					for n := i; n > j; n-- {
						arr[n] = arr[n-1]
					}
					// insert element to the right postion
					arr[j] = current

				}
			}
		}
	}
	sort(originalArr)
	fmt.Println(originalArr)
}
