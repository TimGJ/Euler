package main

import (
	"fmt"
	"github.com/TimGJ/Euler/euler"
)

/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

func AmicableCounterpart(n int) int {
	// Finds if a number is amicable. If it is, returns its counterpart.
	s := euler.SumProperDivisors(n)
	t := euler.SumProperDivisors(s)
	if n == t && n != s {
		fmt.Printf("%d and %d are amicable\n", n, s)
		return s
	}
	return 0
}

func main() {
	var total int
	const limit = 10000
	for i := 0; i < limit; i++ {
		total += AmicableCounterpart(i)
	}
	fmt.Println(total)
}
