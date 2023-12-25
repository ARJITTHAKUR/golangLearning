package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	// fmt.Println("data", openPictureConvertToBytes())
	// os.Exit(1)
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

		parseHTTP(data)
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
	}
}

//	var requestMap = map[string]string{
//		"GET" : ""
//	}
func parseHTTP(data []byte) {
	stringData := string(data)
	// split request body by nextline
	// split by space
	newlineSep := strings.Split(stringData, "\n")
	for _, lines := range newlineSep {
		fmt.Println("value : ", lines)
	}

	// test := []string{"testing", "value", " another "}
	// fmt.Printf("split data: \n %v \n %v", newlineSep, test)
	// fmt.Printf("data recieved :\n %s\n", stringData)
}

func openPictureConvertToBytes() []byte {

	// file, err := os.OpenFile("dog_picture.jpg", os.O_RDONLY, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // pictureData := make([]byte, 2000)
	// var picData []byte

	// file.Read(picData)

	// for {
	// 	var data = make([]byte, 200)
	// 	_, err := file.Readfile(data)
	// 	if err != nil {
	// 		break
	// 	}
	// 	picData = append(picData, data...)
	// }
	// picData = slices.Clip(picData)
	// return picData

	b, err := os.ReadFile("dog_picture.jpg")

	if err != nil {
		log.Fatal(err)
	}
	return b
}
