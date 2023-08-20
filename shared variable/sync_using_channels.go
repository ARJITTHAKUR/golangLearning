package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}
func main() {
	// simplechan := make(chan int)
	for i := 0; i < 5; i++ {
		go func() {
			Deposit(200)
			fmt.Println("=", Balance())
		}()

		go Deposit(400)
	}

	time.Sleep(1 * time.Second)
	// this out of balance is different each time
}
