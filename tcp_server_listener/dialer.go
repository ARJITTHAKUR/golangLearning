package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Println("starting dialer")

	readData := make(chan ([]byte))

	go func(read <-chan []byte) {

		dial, err := net.Dial("tcp", "127.0.0.1:8080")

		if err != nil {
			log.Fatal(err)
		}
		for {
			data := <-read
			fmt.Printf("writing to connection : %s\n", data)
			dial.Write(data)
		}
	}(readData)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		readData <- []byte(line)
	}
}
