package main

import "fmt"

func RemoveItem(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i:= 2
	fmt.Println(RemoveItem(slice, i))
}
