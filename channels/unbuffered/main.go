package main

import "fmt"

func getNaturalNumber(n chan<- int) {
	for i := 0; i < 100; i++ {
		n <- i
	}
	close(n)
}

func multiplyBy2(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * 2
	}
	close(out)
}

func printValues(val <-chan int) {
	for i := range val {
		fmt.Println(i)
	}
}
func main() {

	// channel for natural number
	// channel for multiply by two

	natural := make(chan (int))
	multiply := make(chan (int))

	go getNaturalNumber(natural)
	go multiplyBy2(natural, multiply)
	printValues(multiply)

}
