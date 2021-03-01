package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("go call 1")

	// goroutine with anonymous function
	go func() {
		fun("go call 2")
	}()

	// goroutine with function value call
	fn := fun
	go fn("go call 3")

	// wait for goroutines to end

	// temporarily make main fn sleep so goroutines can be executed - there's a much better way to do this
	fmt.Print("Sleeping \n")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
