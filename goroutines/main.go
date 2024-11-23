package main

import (
	"fmt"
	"sync"
	"time"
)

// A read-write mutex to synchronize access to shared data.
var m = sync.RWMutex{}

// A wait group to ensure the main goroutine waits for all worker goroutines to complete.
var wg = sync.WaitGroup{}

// Simulated database data.
var dbData [5]string = [5]string{"id1", "id2", "id3", "id4", "id5"}

// A slice to store results fetched from the database.
var result []string = []string{}

func main() {
	// Entry point of the program.
	fmt.Println("Goroutines")

	// Record the start time for measuring execution duration.
	t0 := time.Now()

	// Iterate over the database data and start a goroutine for each item.
	for i := range dbData {
		wg.Add(1)    // Increment the WaitGroup counter for each goroutine.
		go dbCall(i) // Start a goroutine to simulate a database call.
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	// Print the total execution time and the resulting data.
	fmt.Println("This is the difference,", time.Since(t0))
	fmt.Println("This is the result data,", result)
}

func dbCall(i int) {
	// Simulates a delay for a database call (fixed 2000ms in this case).
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Save the fetched data to the result slice.
	save(dbData[i])

	// Log the current state of the result slice.
	log()

	// Decrement the WaitGroup counter when the goroutine finishes.
	wg.Done()
}

func save(data string) {
	// Lock the mutex to ensure exclusive access to the shared `result` slice.
	m.Lock()
	result = append(result, data)
	m.Unlock() // Unlock the mutex after the operation is complete.
}

func log() {
	// Acquire a read lock to safely access the `result` slice for logging.
	m.RLock()
	fmt.Printf("\nThe current results are: %v", result)
	m.RUnlock() // Release the read lock after logging.
}
