package main

import (
	"fmt"
	"sync"
)

var (
	counter int // Shared variable accessed by multiple goroutines.
	wg      sync.WaitGroup
)

// increment function, which modifies the shared counter variable.
// This function is intended to be run in a goroutine.
func increment() {
	defer wg.Done() // Decrement the wait group counter on function exit.
	for i := 0; i < 1000; i++ {
		counter++ // Increment the shared counter.
		// Each goroutine attempts to read, modify, and write back the counter value.
		// Since multiple goroutines modify the counter concurrently without synchronization,
		// it leads to a race condition. A race condition occurs when multiple threads or goroutines
		// access a shared variable concurrently, and at least one of the accesses is a write.
		// This can result in an unpredictable outcome because the order of operations is non-deterministic,
		// leading to the final value of counter being dependent on the order in which the goroutines execute.
	}
}

func main() {
	wg.Add(2) // Set the number of goroutines to wait for.

	go increment() // Launch the first goroutine to increment the counter.
	go increment() // Launch the second goroutine to increment the counter.

	wg.Wait()                              // Wait for all goroutines to finish.
	fmt.Println("Final counter:", counter) // Print the final value of the counter.
	// Due to the race condition, this value might not be what is expected (2000 in this case),
	// because the increments may overwrite each other or not be seen by each goroutine.
}
