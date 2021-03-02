package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int
	var wg sync.WaitGroup // define a waitgroup to add deterministic op for concurrent routines

	// Deterministically blocks main thread to allow go routine to be executed.
	wg.Add(1) // add no. of go routines
	go func() {
		defer wg.Done() // decrements no. of goroutines, executed at end of goroutine
		data++
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
