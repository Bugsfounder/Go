package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition - bugsfounder.com")

	wg := &sync.WaitGroup{}
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	wg.Add(3)
	// wg.Add(1)
	// go func(wg *sync.WaitGroup, m *sync.Mutex) {
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	// go func(wg *sync.WaitGroup, m *sync.Mutex) {
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)

}
