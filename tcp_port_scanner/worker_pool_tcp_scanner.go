package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	ports := 1024
	buffer := make(chan int, 10)
	wg := sync.WaitGroup{}

	for i := 0; i < cap(buffer); i++ {
		go worker(&wg, buffer)
	}

	for i := 1; i < ports; i++ {
		wg.Add(1)
		fmt.Println("worker spun:", i)
		buffer <- i
	}

	wg.Wait()
	fmt.Println("closing")
	close(buffer)
}

func worker(wg *sync.WaitGroup, dataChan <-chan int) {
	for data := range dataChan {
		addr := fmt.Sprintf("localhost:%d", data)
		// fmt.Println("scanning", addr)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			writeError(addr)
			continue
		}
		writeLog(addr, data)
		conn.Close()
		wg.Done()
	}

}

func writeLog(addr string, workerIdx int) {
	f, err := os.OpenFile("openPorts.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(fmt.Sprintf("\n worker %d     : %s", workerIdx, addr)))
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
