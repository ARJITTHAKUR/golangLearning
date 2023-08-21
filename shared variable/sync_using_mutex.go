package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	mu      sync.Mutex
	Balance int
}

func (b *Bank) Deposit(amount int) {
	b.mu.Lock()
	b.Balance = b.Balance + amount
	b.mu.Unlock()
}

func (b *Bank) balance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Balance
}

func main() {
	newBank := Bank{
		Balance: 100,
	}

	for i := 0; i < 10; i++ {
		go func(newBank *Bank) {
			// fmt.Println("first")
			newBank.Deposit(100)
			fmt.Println(newBank.balance())
		}(&newBank)

		go func(newBank *Bank) {
			// fmt.Println("second")
			newBank.Deposit(200)
		}(&newBank)
	}

	time.Sleep(time.Second * 1)
}
