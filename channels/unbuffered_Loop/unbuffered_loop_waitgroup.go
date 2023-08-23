package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	dataChan := make(chan int)
	wg := sync.WaitGroup{}

	// wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(i, dataChan, &wg)
	}
	fmt.Println("waiting")
	// close(dataChan)

	// wg.Wait()
	// fmt.Println("closing channel")
	// fmt.Println("returned data :===>", <-dataChan)

	// this doesn't work since it seems like we need to iterate over exact number of goroutine we spun off

	// for {
	// 	data, ok := <-dataChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("data received =>", data)
	// }
	// for data := range dataChan {
	// 	fmt.Println("returned data :===>", data)
	// }

	// solution that works

	for i := 0; i < 10; i++ {
		fmt.Println("waiting")
		fmt.Println("data received ==>", <-dataChan)
		// fmt.Println("")
	}

	close(dataChan)

}

func work(data int, dataChan chan<- int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	// defer wg.Done()
	dataChan <- data
}
