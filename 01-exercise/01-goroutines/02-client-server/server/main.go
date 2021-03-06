package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var i = 0

func main() {
	// TODO: write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "localhost: 12000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		handler, err := listener.Accept()
		if err != nil {
			// if conn failed, go to the next client
			continue
		}

		// handles multiple connections concurrently, Without go routine, it will sequentially work with all connections.
		go handleConn(handler)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server - "+strconv.Itoa(i)+"\n")
		if err != nil {
			return
		}
		i++
		time.Sleep(time.Second)
	}
}
