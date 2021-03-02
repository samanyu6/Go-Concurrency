package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i) //directly accessing i from parent would lead to discrepancy, as parent is continuously changing, therefore,
			// accept i as a parameneter and it'll run with the  value it was passed
		}(i)
	}
	wg.Wait()
}
