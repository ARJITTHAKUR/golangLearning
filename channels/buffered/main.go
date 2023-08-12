package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose       bool
	Cakes         int           // num of cakes to bake
	BakeTime      time.Duration //  time to bake one cake
	BakeStdDev    time.Duration // std deviation of baking time
	BakeBuf       int           // buffer slots b/w baking and icing
	NumIcer       int           // number of cooks doing icing
	IceTime       time.Duration // time to ice one cake
	IceStdDev     time.Duration // std deviation of icing time
	IceBuf        int           // buffer slots b/w icing and incribing
	IncribeTime   time.Duration
	IncribeStdDev time.Duration
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i) // alias for int
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked) // close channel baked after iterating over all the cake
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
	}
	close(iced)
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced

		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.IncribeTime, s.IncribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)

		go s.baker(baked)                // bakes num cake
		for i := 0; i < s.NumIcer; i++ { // ice the cake that is baked
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func work(d, stdDev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stdDev))
	time.Sleep(delay)
}

var defaults = Shop{
	Verbose:       true,
	Cakes:         20,
	BakeTime:      10 * time.Millisecond,
	BakeStdDev:    0,
	BakeBuf:       0,
	NumIcer:       1,
	IceTime:       10 * time.Millisecond,
	IceStdDev:     0,
	IceBuf:        0,
	IncribeTime:   0,
	IncribeStdDev: 0,
}

func main() {
	cakeShop := defaults
	cakeShop.Work(1)
}
