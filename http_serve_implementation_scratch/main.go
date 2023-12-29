package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	list, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening for new requests !")
	for {
		conn, err := list.Accept()
		fmt.Println("got a new request")
		if err != nil {
			fmt.Println(err)
			break
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	data := make([]byte, 1024)
	conn.Read(data)

	// parseHTTP(data)
	// responseData := []byte("HTTP/1.1 200 OK\r\n" +
	// "Accept-Ranges: none\r\n" +
	// "Vary: Accept-Encoding\r\n" +
	// "Content-Type: text/html\r\n" +
	// "Content-lengthL 20\r\n" +
	// "\r\n" +
	// "<h1>Hello, World!")

	responseData := []byte("HTTP/1.1 200 OK\r\n" +
		"Accept-Ranges: none\r\n" +
		"Vary: Accept-Encoding\r\n" +
		"Content-Type: image/jpeg\r\n" +
		// "Content-lengthL 20\r\n" +
		"\r\n")
	responseData = append(responseData, openPictureConvertToBytes()...)
	conn.Write(responseData)
	conn.Close()
}

func parseHTTP(data []byte) {
	stringData := string(data)
	newlineSep := strings.Split(stringData, "\n")
	for _, lines := range newlineSep {
		fmt.Println("value : ", lines)
	}
}

func openPictureConvertToBytes() []byte {

	b, err := os.ReadFile("dog_picture.jpg")

	if err != nil {
		log.Fatal(err)
	}
	return b
}
