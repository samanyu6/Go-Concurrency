package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(5, 4)
	// TODO: get the value computed from goroutine

	r := <-ch
	fmt.Printf("%v", r)
	// fmt.Printf("computed value %v\n", c)
}
