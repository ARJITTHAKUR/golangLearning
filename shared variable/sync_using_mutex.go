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

func (b Bank) Deposit(amount int) {
	b.Balance = b.Balance + amount
}

func (b Bank) balance() int {
	return b.Balance
}

func main() {
	newBank := Bank{
		Balance: 100,
	}

	for i := 0; i < 10; i++ {
		go func(newBank Bank) {
			newBank.Deposit(100)
			fmt.Println(newBank.balance())
		}(newBank)
		go func() {
			newBank.Deposit(200)
		}()
	}

	time.Sleep(time.Second * 1)
}
