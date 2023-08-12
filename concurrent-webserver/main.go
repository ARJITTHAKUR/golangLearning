package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)8

func main() {
	fmt.Println("started")
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // connnetion aborted
			continue
		}
		// handleConn(conn) // handle one connection at a time
		go handleConn(conn) // handles multiple connections at a time

	}
}

func handleConn(c net.Conn) {
	fmt.Println("got a connection")
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
