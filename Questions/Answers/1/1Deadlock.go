package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Goroutine 1: locked the mutex")
		// Simulate some work
		// Trying to lock the mutex again will cause a deadlock
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Goroutine 1: locked the mutex again")
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Goroutine 2: locked the mutex")
	}()

	wg.Wait()
}

//В этом примере Goroutine 1 захватывает мьютекс, а затем пытается захватить его снова, что приводит к deadlock,
//так как мьютекс уже захвачен этой же горутиной и не будет освобожден до завершения работы горутины.
