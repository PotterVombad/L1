package main

import (
	"fmt"
	"sync"
)

func ParallelCycle(array []int) int {
	wait := sync.WaitGroup{}
	ch := make(chan int, 1)
	ch <- 0

	for _, i := range array {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			ch <- i*i + <-ch
		}(i)
	}

	wait.Wait()
	close(ch)

	return <-ch
}

func Conveyor(array []int) int {
	var answer int
	ch := make(chan int)
	go func() {
		for _, val := range array {
			ch <- val * val
		}
		close(ch)
	}()
	for val := range ch {
		answer += val
	}
	return answer
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	fmt.Println(ParallelCycle(array))
	fmt.Println(Conveyor(array))
}
