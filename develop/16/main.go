package main

import (
	"fmt"
)

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left := 0
	base := len(array) - 1

	for i, _ := range array {
		if array[i] < array[base] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	array[left], array[base] = array[base], array[left]

	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}

func main() {
	array := []int{3, 1, 0, 9, 2, 4, 8, 2, 5, 7, 6}
	fmt.Println(QuickSort(array))
}
