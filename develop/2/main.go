package main

import (
	"fmt"
	"sync"
)

func ParallelCycle (array []int) []int{
	answer := make([]int, 0, len(array))
	wait := sync.WaitGroup{}
	ch := make(chan int)

	for _, i := range array {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			ch <- i * i
		}(i)
	}

	go func() {
		wait.Wait()
		close(ch)
	}()

	for val := range ch {
		answer = append(answer, val)
	}
	return answer
}

func Conveyor (array []int) []int{
	answer := make([]int, 0, len(array))
	ch := make(chan int)
	go func() {
		for _, val := range array {
			ch <- val * val
		}
		close(ch)
	}()
	for val := range ch {
		answer = append(answer, val)
	}
	return answer
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	
	fmt.Println(ParallelCycle(array))
	fmt.Println(Conveyor(array))
}
