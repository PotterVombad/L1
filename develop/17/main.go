package main

import (
	"fmt"
)


func binarySearch(array []int, search int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (right + left) / 2

		if array[mid] == search {
			return mid
		} else if array[mid] > search {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	search := 4

	index := binarySearch(array, search)
	if index != -1 {
		fmt.Printf("index of element: %v", index)
	} else {
		fmt.Println("not found")
	}
}
