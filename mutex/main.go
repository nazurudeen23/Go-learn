package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
