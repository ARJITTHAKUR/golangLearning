package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dataChan := make(chan int)

	for i := 0; i < 100; i++ {
		go fetch(i, dataChan)
	}

	go func() {
		for data := range dataChan {
			fmt.Println(data)
		}
	}()
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(<-dataChan)
	// }
	// close(dataChan)

}

func fetch(i int, data chan<- int) {
	x := rand.Intn(100)
	time.Sleep(time.Millisecond * 100 * time.Duration(x))
	data <- i
}
