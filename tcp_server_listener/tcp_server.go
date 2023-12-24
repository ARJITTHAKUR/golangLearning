package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	conn, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TCP server started")

	for {
		fmt.Println("Listening for new connections")
		c, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			for {

				var data = make([]byte, 1024)
				_, err := c.Read(data)
				if err != nil {
					if err == io.EOF {
						fmt.Println("closing connection reached EOF")
						c.Close()
						return
					}
				}
				fmt.Printf("data received : %s\n", string(data))
			}
		}(c)
	}
}
