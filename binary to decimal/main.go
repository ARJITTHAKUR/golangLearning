package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {

	// take an input
	fmt.Println("Please enter an decimal number")

	var userInput string

	_, err := fmt.Scan(&userInput)
	// _, err := fmt.Scan(&userInput)

	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(userInput)

	fmt.Println(userInput)

	if err != nil {
		log.Fatal(err)
	}

	// now convert the n binary digit to decimal number system

	//1101
	var total int
	power := 0
	for n > 0 {
		mod := n % 10
		n = int(n / 10)
		total += (mod * int(math.Exp2(float64(power))))
		power += 1
		// fmt.Println("new val and mod and total", n, mod, total)
	}

	fmt.Println("converted value is ", total)

}
