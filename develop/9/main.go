package main

import (
	"fmt"
)

func Conveyor(array []int) {
	first := make(chan int)
	second := make(chan int)
	go func() {
		for _, val := range array {
			first <- val
		}
		close(first)
	}()
	go func() {
		for val := range first {
			second <- val * val
		}
		close(second)
	}()
	for val := range second {
		fmt.Println(val)
	}
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	Conveyor(array)
}
