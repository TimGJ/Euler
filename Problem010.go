package main

import (
	"fmt"
	"github.com/TimGJ/Euler/euler"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

func main() {
	const max = 2000000

	primes := euler.GetPrimesUpto(max)
	sum := 0
	for _, p := range primes {
		sum += p
	}
	fmt.Println(sum)
}
