package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	bufChan := make(chan string, 2) //sets the limit of buffer of a channel

	for i := 0; i < 10; i++ {
		go fakeWork(i, bufChan) // 10 goroutines are spun off, when buffer of 2 is filled data is sent
	}

	// for data := range bufChan {
	// 	fmt.Println("workDone by : ", data)
	// }

	for i := 0; i < 10; i++ {
		fmt.Println("workDone by : ", <-bufChan)
	}
	fmt.Print("previous work was blocking")
}

func fakeWork(i int, bufchan chan<- string) {
	time.Sleep(time.Second * time.Duration(i))
	bufchan <- strconv.Itoa(i)
}
