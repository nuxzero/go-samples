// Package main demonstrates goroutine communication using channels and synchronization with WaitGroups.
// This program creates multiple reader goroutines that consume values from a shared channel,
// showcasing how Go handles concurrent processing and channel closure detection.
//
// The reader function represents a worker goroutine that:
// - Continuously reads from a receive-only channel
// - Prints received values with its unique ID
// - Gracefully handles channel closure
// - Signals completion via WaitGroup
//
// The main function:
// - Creates a channel and WaitGroup for coordination
// - Launches multiple reader goroutines
// - Sends sequential integers (0-99) to the channel
// - Closes the channel to signal completion
// - Waits for all goroutines to finish
//
// This pattern demonstrates proper channel usage, goroutine synchronization,
// and clean shutdown handling in concurrent Go programs.
//
// Reference https://www.youtube.com/watch?v=nNXhePi3xwE
package main

import (
	"fmt"
	"sync"
)

// reader is a goroutine function that reads integers from a channel.
// It takes an ID for identification, a receive-only channel for reading,
// and a WaitGroup to signal when it has finished processing.
func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	// Defer the Done call to ensure it gets called when the goroutine completes.
	defer wg.Done()
	for {
		// Attempt to read from the channel.
		// If the channel is closed, ok will be false.
		// If ok is false, it means the channel is closed and no more values will be sent.
		// We exit the loop in that case.
		// Otherwise, we print the received value.
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel Closed")
			return
		}
		fmt.Printf("hello random: %d Received: %d\n", id, val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	// Add two to the WaitGroup counter for the two reader goroutines.
	// If you want to add more readers, just increment the count accordingly.
	wg.Add(2)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	// go reader(3, ch, &wg)
	// go reader(4, ch, &wg)

	for i := 0; i < 100; i++ {
		// Send integers 0 to 99 to the channel.
		// The channel is unbuffered, so this will block until a reader is ready
		// to receive the value.
		// If the channel is closed, this will panic, but we handle that in the
		// reader function.
		// If you want to send more values, just change the range limit.
		ch <- i
	}

	// Close the channel to signal that no more values will be sent.
	// This will allow the reader goroutines to exit their loops gracefully.
	// After closing the channel, any further sends to the channel will panic,
	// so ensure that no more values are sent after this point.
	// The readers will detect the closure and exit their loops.
	close(ch)

	// Wait for all reader goroutines to finish processing.
	// This will block until the WaitGroup counter reaches zero,
	// which happens when all reader goroutines call Done.
	// This ensures that the main function does not exit before the readers complete.
	// If you want to add more readers, just ensure to call wg.Add() before starting
	// the goroutines and call wg.Done() in each reader function.
	wg.Wait()
}
