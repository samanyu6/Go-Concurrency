package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

// this is done by moving var i & j from stack to heap so the goroutines still have access post function return

func main() {
	var wg sync.WaitGroup
	var j = 5

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			j++
			fmt.Printf("value of i: %v\n", i)
			fmt.Printf("value of j: %v\n", j)
		}()
		fmt.Println("return from function")
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}
