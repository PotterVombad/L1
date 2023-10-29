package main

import (
	"fmt"
	"sync"
)

// add to map by mutex
func MutexMapAdd(m *sync.Mutex, data map[int]struct{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	data[i] = struct{}{}
	m.Unlock()
}
// sync map
func SyncMapAdd(data *sync.Map, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	data.Store(i, struct{}{})
}

func main() {
	n := 10
	wg := sync.WaitGroup{}
	wg.Add(n)
	mut := sync.Mutex{}
	data := make(map[int]struct{})
	syncData := sync.Map{}

	for i := 0; i < n; i++ {
		go MutexMapAdd(&mut, data, i, &wg)
	}
	wg.Wait()
	fmt.Println("ordenary map")
	for i := range data {
		fmt.Println(i)
	}
	fmt.Println("sync map")
	wg.Add(n)
	for i := 0; i < n; i++ {
		go SyncMapAdd(&syncData, i, &wg)
	}
	wg.Wait()
	syncData.Range(func(k, v interface{}) bool {
		fmt.Println(k)
		return true
	})
}
