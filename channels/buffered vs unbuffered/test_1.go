package main

import (
	"fmt"
	"time"
)

func main() {

	taskchannel := make(chan struct{})
	for i := 0; i < 10; i++ {
		go someTask(taskchannel, i)
	}
	fmt.Println("main thread ")

	// this causes deadlocks
	// for task := range taskchannel {
	// 	fmt.Println("completed task", task)
	// }

	// this works
	for i := 0; i < 10; i++ {
		<-taskchannel
	}

	// but if we do this then the taskchannel is waiting for more data but all the goroutine which is
	// 10 in out case have finished and channel was not closed
	// so main sees that all goroutines are asleep and nothing is left to be done but the channel is still waiting
	// so it throws an error saying all goroutines are asleep

	// for i := 0; i < 20; i++ {
	// 	<-taskchannel
	// }
	// <-taskchannel

	fmt.Println("previous work was blocking")
}

func someTask(taskchannel chan<- struct{}, tasknumber int) {
	fmt.Println("task started", tasknumber)
	time.Sleep(time.Millisecond * 100 * time.Duration(tasknumber))
	fmt.Println("task completed", tasknumber)
	taskchannel <- struct{}{}
}
