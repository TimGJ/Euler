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
	for i := len(d) - 2; i >= 0; i-- {
		s += fmt.Sprintf("%d", d[i])
	}
	return s
}

func main() {

	/* So simulate doing this manually on a piece of paper as 2**1000 is way too large
	for Go integers.
	*/

	// Work out the number of digits required to represent 2**1000

	const exponent = 1000
	width := int(float64(exponent)/math.Log2(10)) + 2
	digits := make(DigitCounters, width)
	carry := make(DigitCounters, width)
	digits[0] = 1
	for i := 1; i <= exponent; i++ {
		for j := 0; j < width-1; j++ {
			digits[j] *= 2
			carry[j+1] += digits[j] / 10
			digits[j] %= 10
		}
		for j := 0; j < width-1; j++ {
			digits[j] += carry[j]
			carry[j] = 0
		}
	}
	var digitsum int
	for _, val := range digits {
		digitsum += val
	}
	fmt.Printf("2**%d=%s. Digit sum=%d\n", exponent, digits, digitsum)
}
