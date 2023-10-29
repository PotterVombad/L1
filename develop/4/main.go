package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(wg *sync.WaitGroup,id int, ch chan int) {
	for val := range ch {
		fmt.Printf("worker id: %v, inf: %v \n", id, val)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()

}

func main() {
	var N int = 2
	count, err := fmt.Scan(&N)
	if err != nil || N < 1 || count != 1 {
		fmt.Println("You need to write a number starting from 1")
		return
	}

	var wg sync.WaitGroup
	wg.Add(N)
	ch := make(chan int)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for id := 0; id < N; id++ {
		go Worker(&wg, id, ch)
	}

loop:
	for {
		select {
		case <-sigs:
			close(ch)
			break loop
		default:
			data := rand.Int()
			ch <- data
		}
	}

	wg.Wait()
	fmt.Println("you stop workers")
}
