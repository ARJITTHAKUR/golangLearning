package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Commencing countdown...")

	tick := time.Tick(time.Second * 1)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //read single byte
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("launch abort")
			return
			// default: // created non blocking behaviour and loop runs without blocking and listening for events from any of the above chan

		}

	}

	launch()
}

func launch() {
	fmt.Println("launched")
}
