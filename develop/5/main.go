package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func Writer(dur int, ch chan int, wait *sync.WaitGroup) () {
	duration := time.Millisecond * time.Duration(dur)
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			close(ch)
			return
		default:
			ch <- rand.Int()
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func Reader(ch chan int, wait *sync.WaitGroup) () {
	for val := range ch {
		fmt.Println(val)
	}
	wait.Done()
}

func main() {
	fmt.Println("Введите число")
	var time int 
	scan, err := fmt.Scan(&time)
	if err != nil || scan != 1 || time <= 0 {
		panic("argument is not good")
	}
	ch := make(chan int)
	var wait sync.WaitGroup
	go Writer(time, ch, &wait)
	go Reader(ch, &wait)
	wait.Add(1)
	wait.Wait()
	fmt.Println("Время вышло")
}