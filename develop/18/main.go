package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
// сделано две реализации с mutex и atomic
type (
	CounterMutex struct {
		count int
		mut   sync.Mutex
	}
	CounterAtomic struct {
		count atomic.Int64
	}
)

func NewCounterMutex() *CounterMutex {
	return &CounterMutex{
		count: 0,
		mut:   sync.Mutex{},
	}
}

func (cm *CounterMutex) IncreamentCM() {
	cm.mut.Lock()
	cm.count++
	cm.mut.Unlock()
}

func NewCounterAtomic() *CounterAtomic {
	return &CounterAtomic{
		count: atomic.Int64{},
	}
}

func (ca *CounterAtomic) IncreamentCA() {
	ca.count.Add(1)
}

func main() {
	n := 100
	cm := NewCounterMutex()
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cm.IncreamentCM()
		}()
	}
	wg.Wait()
	fmt.Printf("count with mutex, %v",cm.count)
	ca := NewCounterAtomic()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ca.IncreamentCA()
		}()
	}
	wg.Wait()
	fmt.Printf("count with atomic, %v",ca.count.Load())

}
