package main

import (
	"fmt"
	"math"
)

/*
Power digit sum
Problem 16
2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2**1000?
*/

type DigitCounters []int

func (d DigitCounters) String() string {
	var s string
	for i := len(d) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%d", d[i])
	}
	return s
}

func main() {

	/* So simulate doing this manually on a piece of paper as 2**1000 is way too large
	for Go integers.
	*/

	// Work out the number of digits required to represent 2**1000

	const exponent = 10
	width := int(float64(exponent)/math.Log2(10)) + 1
	digits := make(DigitCounters, width)
	carry := make(DigitCounters, width)
	digits[0] = 1
	fmt.Println(digits)
	for i := 1; i <= exponent; i++ {
		for j := 0; j < width-1; j++ {
			digits[j] *= 2
			carry[j] = digits[j] / 10
			digits[j] %= 10
			digits[j+1] += carry[j]
		}
		fmt.Printf("2 ** %d = %s\n", i, digits)
	}

}
