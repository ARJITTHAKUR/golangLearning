package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	list, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("accepting requests !")
		conn, err := list.Accept()
		fmt.Println("got a new request")
		if err != nil {
			fmt.Println(err)
			break
		}
		data := make([]byte, 1024)
		conn.Read(data)

		fmt.Printf("data recieved :\n %s\n", data)

	}
}
