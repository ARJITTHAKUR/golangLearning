package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// take an input
	fmt.Println("Please enter an decimal number")

	var userInput string

	_, err := fmt.Scan(&userInput)

	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(userInput)

	fmt.Println(userInput)

	if err != nil {
		log.Fatal(err)
	}

	// now convert the n binary digit to decimal number system

	// var total int
	// power := 0
	// for n > 0 {
	// 	mod := n % 10
	// 	n = int(n / 10)
	// 	total += (mod * int(math.Exp2(float64(power))))
	// 	power += 1
	// 	// fmt.Println("new val and mod and total", n, mod, total)
	// }

	total, err := binaryToDecimal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("converted value is ", total)

}
func decimalToBinary(num int) {

}
func binaryToDecimal(num int) (int, error) {

	binary := ""
	for num > 0 {
		temp := num
		times := 0
		for temp > 1 {
			temp = temp - 2
			times += 1
			// fmt.Println(temp)

		}
		fmt.Println(times)
		remainder := num % 2
		num = times
		binary = fmt.Sprintf("%s%d", binary, remainder)

		// binary = binary + remainder
		// binary = fmt.Sprintf("%s%d", binary, remainder)
		// fmt.Println(remainder, num, binary)
	}

	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}
	binary = reverse(binary)
	interger, err := strconv.Atoi(binary)

	if err != nil {
		return 0, err
	}

	return interger, nil
}
