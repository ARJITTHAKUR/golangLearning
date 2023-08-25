package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	for i := 0; i <= 65535; i++ {
		addr := fmt.Sprintf("localhost:%d", i)
		fmt.Println("scanning", addr)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			writeError(addr)
			continue
		}
		writeLog(addr)
		conn.Close()
	}
}

func writeLog(addr string) {
	f, err := os.OpenFile("openPorts.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(fmt.Sprintf("\n %s", addr)))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func writeError(addr string) {
	f, err := os.OpenFile("errors.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(fmt.Sprintf("\n %s", addr)))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
