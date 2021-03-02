package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch) // thread safe exit
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			ch <- i
		}
	}()

	// TODO: range over channel to recv values
	for x := range ch {
		fmt.Println(x)
	}
}
