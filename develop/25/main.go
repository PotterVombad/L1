package main

import (
	"context"
	"fmt"
	"time"
)

func SleepTimeAfter(sleep time.Duration) {
	<- time.After(sleep)
	fmt.Printf("time after SleepTimeAfter: %v \n", time.Now())

}

func SleepTimeTick(sleep time.Duration) {
	<- time.Tick(sleep)
	fmt.Printf("time after SleepTimeTick: %v \n", time.Now())
}

func SleepContext(ctx context.Context, sleep time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, sleep)
	defer cancel() 
	<- ctx.Done()
	fmt.Printf("time after SleepContext: %v \n", time.Now())
}

func main() {
	n := "5s"
	sleep, err := time.ParseDuration(n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("massege before SleepTimeAfter: %v \n", time.Now())
	SleepTimeAfter(sleep)
	fmt.Printf("massege before SleepTimeTick: %v \n", time.Now())
	SleepTimeTick(sleep)
	ctx := context.Background()
	fmt.Printf("massege before SleepContext: %v \n", time.Now())
	SleepContext(ctx, sleep)
}
