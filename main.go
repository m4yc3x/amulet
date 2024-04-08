package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting Greyrose server...")

	var wg sync.WaitGroup

	// Increment the WaitGroup counter for each goroutine
	wg.Add(3)

	// Starting server tasks
	go LoginServerRoutine(&wg)
	go PatchServerRoutine(&wg)
	go GameServerRoutine(&wg)

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Printf("SERVER CLOSED @ %s\n", time.Now().Format(time.RFC1123))
}
