package main

import (
	"fmt"
	"sync"
	"time"
)

func semaforStop(wg *sync.WaitGroup, ch <-chan struct{}) {
	wg.Add(1)
	for {
		select {
		case <-ch:
			wg.Done()
			return
		default:
			fmt.Println("not yet")
			time.Sleep(time.Second)
		}
	}
}

func timeStop(wg *sync.WaitGroup, t int) {
	duration := time.Second * time.Duration(t)
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			wg.Done()
			return
		default:
			fmt.Println("not yet")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	go semaforStop(&wg, ch)
	time.Sleep(5 * time.Second)
	ch <- struct{}{}
	wg.Wait()
	fmt.Println("semafor done")

	wg.Add(1)
	duration := 5
	go timeStop(&wg, duration)
	wg.Wait()
	fmt.Println("timer done")
}
