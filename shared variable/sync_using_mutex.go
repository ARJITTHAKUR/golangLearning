package main

import (
	"fmt"
	"time"
)

type Bank struct {
	// mu      sync.Mutex
	Balance int
}

func (b *Bank) Deposit(amount int) {
	fmt.Println("despositing start")

	// b.mu.Lock()
	b.Balance = b.Balance + amount
	// time.Sleep(500 * time.Millisecond)
	fmt.Println("despositing completed")
	// b.mu.Unlock()
}

func (b *Bank) balance() int {
	fmt.Println("balance")
	// b.mu.Lock()
	// defer b.mu.Unlock()
	// time.Sleep(500 * time.Millisecond)
	fmt.Println("balance complete")
	return b.Balance
}

func main() {
	newBank := Bank{
		Balance: 100,
	}

	for i := 0; i < 5; i++ {
		// deposites
		go func(newBank *Bank) {
			// fmt.Println("first")
			newBank.Deposit(100)
			// fmt.Println(newBank.balance())
		}(&newBank)

		// reads
		go func(newBank *Bank) {
			fmt.Println(newBank.balance())
		}(&newBank)
		// deposites

		go func(newBank *Bank) {
			// fmt.Println("second")
			newBank.Deposit(200)
		}(&newBank)
	}
	fmt.Println("final balance", newBank.balance())
	time.Sleep(time.Second * 10)
}
