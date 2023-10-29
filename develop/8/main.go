package main

import "fmt"

func setBit(num int64, index int) int64 {
	value := num & (1 << index)
	if value > 0 {
		return num &^ (1 << uint(index))
	} else {
		return num | (1 << uint(index))
	}
}

func main() {
	var num int64 = 64 
	index := 1      

	result := setBit(num, index)
	fmt.Printf("Original number: %d, in bytes: %08b\n", num, num)
	fmt.Printf("Result number: %d, in bytes: %08b", result, result)
}
