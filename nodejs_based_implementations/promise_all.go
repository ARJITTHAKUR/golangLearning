package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int, 10)
	// wg := sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		// wg.Add(1)
		go fetch(i, dataChan)
	}
	// for data := range dataChan {
	// 	fmt.Println(data)
	// }
	for i := 0; i < 200; i++ {
		fmt.Println(<-dataChan)
	}
	// wg.Wait()
	// time.Sleep(1 * time.Second)
}

func fetch(i int, data chan<- int) {
	// defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	data <- i
}
