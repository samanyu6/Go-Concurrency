package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(rec chan string) {
	// send message on ch1
	rec <- "yo"
}

func relayMsg(rec <-chan string, send chan<- string) {
	// recv message on ch1
	// send it on ch2
	temp := <-rec
	send <- temp
}

func main() {
	// create ch1 and ch2
	rec := make(chan string)
	send := make(chan string)
	// spine goroutine genMsg and relayMsg

	go genMsg(rec)

	go relayMsg(rec, send)
	// recv message on ch2

	ans := <-send
	fmt.Println("Ans ", ans)
}
